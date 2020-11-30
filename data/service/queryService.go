package service

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/qxip/cloki-go/model"
	"github.com/qxip/cloki-go/utils/heputils"
	"github.com/sirupsen/logrus"
)

type QueryService struct {
	ServiceData
}

type QueryStruct struct {
	DB       string
	Table    string
	Interval string
	Tag      string
	Metric   string
	Where    string
}

// this method create new user in the database
// it doesn't check internally whether all the validation are applied or not
func (qs *QueryService) GetQuery() string {
	reply := gabs.New()
	reply.Set("success", "status")
	reply.Set("streams", "data", "resultType")
	reply.ArrayAppend([]string{}, "data", "result")
	return reply.String()
}

// this method create new user in the database
// it doesn't check internally whether all the validation are applied or not
func (qs *QueryService) clickHouseQueryData(query string, tags string) string {

	//timeSeries := []model.TableTimeSeries{}
	rows, err := qs.Session.Queryx(query) // (*sql.Rows, error)
	if err != nil {
		logrus.Error("couldn't query data: ", err.Error())

	}
	var objects []map[string]interface{}

	for rows.Next() {

		columns, err := rows.ColumnTypes()

		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}

		for i, column := range columns {
			v := reflect.New(column.ScanType()).Interface()
			switch k := v.(type) {
			case *string:
				values[i] = k
			case *[][]interface{}:
				values[i] = k
			default:
				values[i] = k
			}

			object[column.Name()] = v
		}

		err = rows.Scan(values...)
		if err != nil {
			logrus.Error("Error while selecting values", err)
		}
		objects = append(objects, object)
	}
	data, err := json.Marshal(objects)
	if err != nil {
		logrus.Error(err)
	}

	tagsArr := strings.Split(tags, ",")
	reply := gabs.New()
	reply.Set("success", "status")
	reply.Set("matrix", "data", "resultType")
	result, _ := gabs.ParseJSON([]byte(data))
	for _, metric := range result.Children() {
		st := gabs.New()
		lv := gabs.New()
		if metric.Exists("groupArr") {
			groupArr := metric.Search("groupArr").Data().([]interface{})
			lv.ArrayOfSize(len(groupArr), "values")
			for k, m := range metric.Search("groupArr").Data().([]interface{}) {
				var values []interface{}
				gData := m.([]interface{})
				values = append(values, gData[0].(float64)/1000)
				values = append(values, fmt.Sprintf("%g", gData[1]))
				lv.S("values").SetIndex(values, k)
			}
		}
		for _, tg := range tagsArr {
			if tg != "groupArr" {
				lb := strings.TrimSpace(tg)
				st.Set(metric.Search(lb).Data(), "metric", lb)
			}
		}
		lv.Merge(st)
		reply.ArrayAppend(lv, "data", "result")
	}

	if !reply.Exists("result"){
		reply.Set([]string{}, "data", "result")
	}
	return reply.String()
	return string(data)
}

// this method create new user in the database
// it doesn't check internally whether all the validation are applied or not
func (qs *QueryService) QueryRange(query, start, end string) string {

	rateQuery := `(.*) by \((.*)\) \(rate\((.*)\[(.*)\]\)\) from (.*)\.(.*)`
	rateQueryWhere := `(.*) by \((.*)\) \(rate\((.*)\[(.*)\]\)\) from (.*)\.(.*) where (.*)`
	rateQueryNoWhere := `(.*) by \((.*)\) \(rate\((.*)\[(.*)\]\)\) from (.*)\.([\S]+)\s?(.*)`

	var template string
	if match, _ := regexp.MatchString(rateQuery, query); match {
		re := regexp.MustCompile(rateQuery)
		s := re.FindAllStringSubmatch(query, -1)[0]
		settings := QueryStruct{
			DB:       s[5],
			Table:    s[6],
			Interval: s[4],
			Tag:      s[2],
			Metric:   fmt.Sprintf("%s(%s)", s[1], s[3]),
			Where:    "",
		}
		// Lets query!
		template = fmt.Sprintf("SELECT %s, groupArray((t, c)) AS groupArr FROM (SELECT (intDiv(toUInt32(%s), "+
			"%s) * %s) * 1000 AS t, %s, %s c FROM %s.%s", settings.Tag, "record_datetime", settings.Interval,
			settings.Interval, settings.Tag, settings.Metric, settings.DB, settings.Table)

		if start != "" && end != "" {
			template += fmt.Sprintf(" PREWHERE record_datetime BETWEEN %d AND %d", heputils.GetSeconds(start), heputils.GetSeconds(end))
		}

		if settings.Where != "" {
			template += " AND " + settings.Where
		}
		template += fmt.Sprintf(" GROUP BY t, %s ORDER BY t, %s) GROUP BY %s ORDER BY %s", settings.Tag, settings.Tag, settings.Tag, settings.Tag)
		logrus.Debug(template)
		return qs.clickHouseQueryData(template, settings.Tag)
	}

	if match, _ := regexp.MatchString(rateQueryWhere, query); match {
		re := regexp.MustCompile(rateQueryWhere)
		s := re.FindAllStringSubmatch(query, -1)[0]
		settings := QueryStruct{
			DB:       s[5],
			Table:    s[6],
			Interval: s[4],
			Tag:      s[2],
			Metric:   fmt.Sprintf("%s(%s)", s[1], s[3]),
			Where:    s[7],
		}
		// Lets query!
		template = fmt.Sprintf("SELECT %s, groupArray((t, c)) AS groupArr FROM (SELECT (intDiv(toUInt32(%s), "+
			"%s) * %s) * 1000 AS t, %s, %s c FROM %s.%s", settings.Tag, "record_datetime", settings.Interval,
			settings.Interval, settings.Tag, settings.Metric, settings.DB, settings.Table)

		if start != "" && end != "" {
			template += fmt.Sprintf(" PREWHERE record_datetime BETWEEN %d AND %d", heputils.GetSeconds(start), heputils.GetSeconds(end))
		}

		if settings.Where != "" {
			template += " AND " + settings.Where
		}
		template += fmt.Sprintf(" GROUP BY t, %s ORDER BY t, %s) GROUP BY %s ORDER BY %s", settings.Tag, settings.Tag, settings.Tag, settings.Tag)
		logrus.Debug(template)
		return qs.clickHouseQueryData(template, settings.Tag)
	}

	if strings.HasPrefix(query, "clickhouse"){
		query = strings.TrimPrefix(query, "clickhouse")
		query = query[2:len(query)-2]
		var settings QueryStruct
		for _, value := range strings.Split(query, ","){
			fmt.Println(value)
			if strings.Contains(value, "db="){
				settings.DB = strings.ReplaceAll(strings.Split(value, "=")[1], "\"", "")
			}
			if strings.Contains(value, "table="){
				settings.Table = strings.ReplaceAll(strings.Split(value, "=")[1], "\"", "")
			}
			if strings.Contains(value, "tag="){
				settings.Tag = strings.ReplaceAll(strings.Split(value, "=")[1], "\"", "")
			}
			if strings.Contains(value, "metric="){
				settings.Metric = strings.ReplaceAll(strings.Split(value, "=")[1], "\"", "")
			}
			if strings.Contains(value, "interval="){
				settings.Interval = strings.ReplaceAll(strings.Split(value, "=")[1], "\"", "")
			}
		}

		if settings.Interval == ""{
			settings.Interval = "60"
		}

		// Lets query!
		template = fmt.Sprintf("SELECT %s, groupArray((t, c)) AS groupArr FROM (SELECT (intDiv(toUInt32(%s), "+
			"%s) * %s) * 1000 AS t, %s, %s c FROM %s.%s", settings.Tag, "record_datetime", settings.Interval,
			settings.Interval, settings.Tag, settings.Metric, settings.DB, settings.Table)

		if start != "" && end != "" {
			template += fmt.Sprintf(" PREWHERE record_datetime BETWEEN %d AND %d", heputils.GetSeconds(start), heputils.GetSeconds(end))
		}

		if settings.Where != "" {
			template += " AND " + settings.Where
		}
		template += fmt.Sprintf(" GROUP BY t, %s ORDER BY t, %s) GROUP BY %s ORDER BY %s", settings.Tag, settings.Tag, settings.Tag, settings.Tag)
		logrus.Debug(template)
		return qs.clickHouseQueryData(template, settings.Tag)
	}

	if match, _ := regexp.MatchString(rateQueryNoWhere, query); match {
		re := regexp.MustCompile(rateQueryNoWhere)
		s := re.FindAllStringSubmatch(query, -1)[0]
		settings := QueryStruct{
			DB:       s[5],
			Table:    s[6],
			Interval: s[4],
			Tag:      s[2],
			Metric:   fmt.Sprintf("%s(%s)", s[1], s[3]),
			Where:    s[3],
		}
		// Lets query!
		template = fmt.Sprintf("SELECT %s, groupArray((t, c)) AS groupArr FROM (SELECT (intDiv(toUInt32(%s), "+
			"%s) * %s) * 1000 AS t, %s, %s c FROM %s.%s", settings.Tag, "record_datetime", settings.Interval,
			settings.Interval, settings.Tag, settings.Metric, settings.DB, settings.Table)

		if start != "" && end != "" {
			template += fmt.Sprintf(" PREWHERE record_datetime BETWEEN %d AND %d", heputils.GetSeconds(start), heputils.GetSeconds(end))
		}

		if settings.Where != "" {
			template += " AND " + settings.Where
		}
		template += fmt.Sprintf(" GROUP BY t, %s ORDER BY t, %s) GROUP BY %s ORDER BY %s", settings.Tag, settings.Tag, settings.Tag, settings.Tag)
		logrus.Debug(template)
		return qs.clickHouseQueryData(template, settings.Tag)
	}

	re := regexp.MustCompile(`\{\"?\b(\w+)\"?(!?=~?)\"?([^\"\\n]*?)\"`)
	prine := re.FindAllSubmatch([]byte(query), -1)
	var labelRules []model.LabelRules
	for _, i := range prine {
		var lr model.LabelRules
		lr.Label = string(i[1])
		lr.Cond = string(i[2])
		lr.Value = string(i[3])
		labelRules = append(labelRules, lr)

	}
	var condition []string
	for _, lr := range labelRules {
		switch lr.Cond {
		case "=":
			condition = append(condition, fmt.Sprintf("(visitParamExtractString(labels, '%s') = '%s')", lr.Label, lr.Value))
		case "!=":
			condition = append(condition, fmt.Sprintf("(visitParamExtractString(labels, '%s') != '%s')", lr.Label, lr.Value))
		}
	}
	var fingerSearch = "SELECT DISTINCT fingerprint FROM time_series FINAL PREWHERE " + strings.Join(condition, "OR")

	//timeSeries := []model.TableTimeSeries{}
	rows, err := qs.Session.Queryx(fingerSearch) // (*sql.Rows, error)
	if err != nil {
		logrus.Error("couldn't query data: ", err.Error())

	}
	defer rows.Close()

	var fingerPrints []uint64
	for rows.Next() {
		var finerprint uint64
		rows.Scan(&finerprint)
		fingerPrints = append(fingerPrints, finerprint)

	}

	var selectQuery = "SELECT fingerprint, timestamp_ms, string" + " FROM samples" + " WHERE fingerprint IN (" +
		strings.Trim(strings.Join(strings.Fields(fmt.Sprint(fingerPrints)), ","), "[]") + ")"

	if len(start) != 0 && len(end) != 0 {
		selectQuery += fmt.Sprintf(" AND timestamp_ms BETWEEN %d AND %d", heputils.PartTime(start), heputils.PartTime(end))
	}

	selectQuery += " ORDER BY fingerprint, timestamp_ms"

	logrus.Debug(selectQuery)

	//timeSeries := []model.TableTimeSeries{}
	rows, err = qs.Session.Queryx(selectQuery) // (*sql.Rows, error)
	if err != nil {
		logrus.Error("couldn't query data: ", err.Error())

	}

	defer rows.Close()

	var samples []model.TableSample
	for rows.Next() {
		var sp model.TableSample
		rows.Scan(&sp.FingerPrint, &sp.TimestampMS, &sp.String)
		samples = append(samples, sp)

	}

	lv := gabs.New()
	lv.ArrayOfSize(len(samples), "values")

	var values []string
	for k, sp := range samples {
		values = append(values, strconv.FormatInt(int64(time.Duration(sp.TimestampMS)*time.Millisecond/time.Nanosecond), 10), sp.String)
		lv.S("values").SetIndex(values, k)

	}

	st := gabs.New()
	st.Set(labelRules[0].Value, "stream", labelRules[0].Label)

	lv.Merge(st)

	reply := gabs.New()
	reply.Array("data", "result")
	reply.ArrayAppend(lv, "data", "result")

	reply.Set("streams", "data", "resultType")

	return reply.String()
}
