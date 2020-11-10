package service

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/sirupsen/logrus"
	"gitlab.com/qxip/cloki/model"
	"gitlab.com/qxip/cloki/utils/heputils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type QueryService struct {
	ServiceData
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
func (qs *QueryService) QueryRange(query, start, end string) string {

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
		logrus.Error("couldn't select alias data: ", err.Error())

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
		logrus.Error("couldn't select alias data: ", err.Error())

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
