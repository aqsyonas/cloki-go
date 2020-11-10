package service

import (
	"database/sql"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"github.com/valyala/bytebufferpool"
	"gitlab.com/qxip/cloki/model"
	"gitlab.com/qxip/cloki/utils/heputils"
	"gitlab.com/qxip/cloki/utils/logger/function"
	"sort"
	"strconv"
	"strings"
	"time"
)

type PushService struct {
	ServiceData
	GoCache *cache.Cache
	TSCh    chan *model.TableTimeSeries
	SPCh    chan *model.TableSample
	DBTimer time.Duration
	DBBulk  int
}

func (ps *PushService) Insert() error {

	timer := time.NewTimer(ps.DBTimer)
	stop := func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
	}
	defer stop()

	var (
		txTS *sql.Tx
		txSP *sql.Tx
	)
	var (
		stmtTS *sql.Stmt
		stmtSP *sql.Stmt
	)

	var (
		tsCnt int
		spCnt int
	)
	var err error

	sqlTS := fmt.Sprintf("INSERT INTO time_series (%s) VALUES (%s)",
		function.FieldName(function.DBFields(model.TableTimeSeries{})), function.FieldValue(function.DBFields(model.TableTimeSeries{})))

	sqlSP := fmt.Sprintf("INSERT INTO samples (%s) VALUES (%s)",
		function.FieldName(function.DBFields(model.TableSample{})), function.FieldValue(function.DBFields(model.TableSample{})))

	for {
		select {
		case ts, _ := <-ps.TSCh:
			if tsCnt == 0 {
				txTS, err = ps.Session.Begin()
				if err != nil {
					logrus.Fatal(err)
				}
				stmtTS, err = txTS.Prepare(sqlTS)
				if err != nil {
					logrus.Fatal(err)
				}
			}
			stmtTS.Exec(function.GenerateArg(ts)...)
			tsCnt++

		case sample, _ := <-ps.SPCh:
			if spCnt == 0 {
				txSP, err = ps.Session.Begin()
				if err != nil {
					logrus.Fatal(err)
				}
				stmtSP, err = txSP.Prepare(sqlSP)
				if err != nil {
					logrus.Fatal(err)
				}
			}
			stmtSP.Exec(function.GenerateArg(sample)...)
			spCnt++
			if spCnt == ps.DBBulk {
				err := txSP.Commit()
				if err != nil {
					logrus.Error(err)
				}
				err = txTS.Commit()
				if err != nil {
					logrus.Error(err)
				}
				tsCnt = 0
				spCnt = 0
			}

		case <-timer.C:
			timer.Reset(ps.DBTimer)
			switch {
			case tsCnt > 0:
				err := txTS.Commit()
				if err != nil {
					logrus.Error(err)
				}
				tsCnt = 0
			case spCnt > 0:
				err = txSP.Commit()
				if err != nil {
					logrus.Error(err)
				}
				spCnt = 0
			}
		}
	}
}

// this method create new user in the database
// it doesn't check internally whether all the validation are applied or not
func (ps *PushService) PushStream(req model.PushRequest) error {

	for _, stream := range req.Streams {
		labelsArr := strings.Split(stream.Labels[1:len(stream.Labels)-1], ",")

		lbs := make([]model.Label, len(labelsArr))
		labelKey := make([]string, len(labelsArr))

		labelValue := make(map[string][]string)
		for k, l := range labelsArr {
			keyValue := strings.SplitN(l, "=", 2)
			value, _ := strconv.Unquote(keyValue[1])
			labelKey[k] = keyValue[0]
			labelValue[keyValue[0]] = append(labelValue[keyValue[0]], value)
			lbs[k] = model.Label{
				Key:   keyValue[0],
				Value: value,
			}
		}

		// lets insert only the unique values for key
		for k, v := range labelValue {
			if keys, exist := ps.GoCache.Get(k); exist {
				ps.GoCache.Replace(k, heputils.AppendTwoSlices(keys.([]string), heputils.UniqueSlice(v)), 0)
			} else {
				ps.GoCache.Add(k, heputils.UniqueSlice(v), 0)
			}
		}

		sort.Slice(lbs[:], func(i, j int) bool {
			return lbs[i].Key < lbs[j].Key
		})

		fingerPrint := heputils.FingerprintLabels(lbs)

		// if fingerprint was not found, lets insert into time_series
		if _, found := ps.GoCache.Get(fmt.Sprint(fingerPrint)); !found {
			if keys, exist := ps.GoCache.Get("__LABEL__"); exist {
				labelKeys := keys.([]string)
				uniqueKeys := heputils.AppendTwoSlices(labelKeys, labelKey)
				ps.GoCache.Replace("__LABEL__", uniqueKeys, 0)

			} else {
				ps.GoCache.Add("__LABEL__", labelKey, 0)
			}

			b := bytebufferpool.Get()

			ps.GoCache.Set(fmt.Sprint(fingerPrint), true, cache.DefaultExpiration)

			ps.TSCh <- &model.TableTimeSeries{
				Date:        time.Now(),
				FingerPrint: fingerPrint,
				Labels:      heputils.MakeJson(lbs, b),
				Name:        "",
			}

		}

		for _, entries := range stream.Entries {

			ps.SPCh <- &model.TableSample{
				FingerPrint: fingerPrint,
				TimestampMS: time.Now().UnixNano() / 1000000,
				Value:       0,
				String:      entries.Line,
			}

		}
	}
	return nil
}

// this method create new user in the database
// it doesn't check internally whether all the validation are applied or not
func (ps *PushService) ReloadFingerprints() error {

	//timeSeries := []model.TableTimeSeries{}
	rows, err := ps.Session.Queryx("SELECT DISTINCT fingerprint, labels FROM time_series") // (*sql.Rows, error)
	if err != nil {
		logrus.Error("couldn't select alias data: ", err.Error())
	}

	defer rows.Close()
	var labels []string
	for rows.Next() {
		var label string
		var finerprint uint64
		rows.Scan(&finerprint, &label)
		labels = append(labels, label)

	}

	for _, label := range labels {
		lb, _ := gabs.ParseJSON([]byte(label))
		var labelKey []string
		labelValue := make(map[string][]string)
		for lk, lv := range lb.ChildrenMap() {
			labelKey = append(labelKey, lk)
			labelValue[lk] = append(labelValue[lk], lv.Data().(string))
		}

		// lets have only unique values for a label keys
		for k, v := range labelValue {
			if keys, exist := ps.GoCache.Get(k); exist {
				ps.GoCache.Replace(k, heputils.AppendTwoSlices(keys.([]string), heputils.UniqueSlice(v)), 0)
			} else {
				ps.GoCache.Add(k, heputils.UniqueSlice(v), 0)
			}
		}

		// lets have only unique label keys
		if keys, exist := ps.GoCache.Get("__LABEL__"); exist {
			labelKeys := keys.([]string)
			uniqueKeys := heputils.AppendTwoSlices(labelKeys, labelKey)
			ps.GoCache.Replace("__LABEL__", uniqueKeys, 0)

		} else {
			ps.GoCache.Add("__LABEL__", labelKey, 0)
		}

	}
	return nil
}
