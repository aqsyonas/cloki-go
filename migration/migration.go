package migration

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/qxip/cloki-go/model"
	"github.com/qxip/cloki-go/utils/heputils"
)

type RollesTable struct {
	Username   string `json:"Username"`
	Attributes string `json:"Attributes"`
}

// getSession creates a new root session and panics if connection error occurs
func GetDataRootDBSession(user *string, password *string, dbname *string, host *string, port *int) (*sqlx.DB, error) {

	if *port != 0 {
		*port = 9000
	}

	connectString := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s&read_timeout=%d&write_timeout=%d&compress=true&debug=true",
		*host, *port, *user, *password, *dbname, 10, 30)

	db, err := sqlx.Open("clickhouse", connectString)

	heputils.Colorize(heputils.ColorYellow, fmt.Sprintf("\nCONNECT to DB ROOT STRING: [%s]\n", connectString))

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {

		if exception, ok := err.(*clickhouse.Exception); ok {
			logrus.Error(fmt.Sprintf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace))
		} else {
			logrus.Info("ping db data ", err)
		}

		return nil, nil
	}

	logrus.Info("----------------------------------- ")
	logrus.Info("*** Database Data Root Session created *** ")
	logrus.Info("----------------------------------- ")
	return db, nil
}

func CreateNewUser(dataRootDBSession *sqlx.DB, user *string, password *string) {

	createString := fmt.Sprintf("\r\nCLoki - creating user [user=%s password=%s]", *user, *password)

	heputils.Colorize(heputils.ColorRed, createString)

	sql := fmt.Sprintf("CREATE USER %s WITH PASSWORD  '%s'", *user, *password)

	dataRootDBSession.Exec(sql)

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")

}

func CreateClokiDB(dataRootDBSession *sqlx.DB, dbname *string, user *string) {

	createString := fmt.Sprintf("\r\nCLoki - create db [%s] with [name=%s]", *dbname, *user)

	heputils.Colorize(heputils.ColorRed, createString)

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s ", *dbname)

	dataRootDBSession.Exec(sql)

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

func CheckVersion(dataRootDBSession *sqlx.DB) (int, error) {

	var VersionNum int = 20

	return VersionNum, nil
}

func CreateClokiTables(DBSession *sqlx.DB, clokiDBName string) {

	createString := fmt.Sprintf("\r\nCLoki - creating tables for the Cloki DB [dbname=%s]", clokiDBName)
	heputils.Colorize(heputils.ColorGreen, createString)

	var executeSql []string

	sample := &model.TableSample{}
	sql := GenerateSQL(sample.TableName(), sample, sample.TableEngine())
	executeSql = append(executeSql, sql)

	timeSeries := &model.TableTimeSeries{}
	sql = GenerateSQL(timeSeries.TableName(), timeSeries, timeSeries.TableEngine())
	executeSql = append(executeSql, sql)

	for _, val := range executeSql {
		_, err := DBSession.Exec(val)
		if err != nil {
			logrus.Error(fmt.Sprintf("Automigrate failed: with error %s", err.Error))
		}
	}

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

// DBFields reflects on a struct and returns the values of fields with `db` tags,
// or a map[string]interface{} and returns the keys.
func GenerateSQL(table string, values interface{}, engine string) string {

	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	sql := "CREATE TABLE IF NOT EXISTS " + table + " (\n"
	masterKey := ""
	fields := []string{}
	orders := []string{}

	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" && field != "-" {
				clickhouse := v.Type().Field(i).Tag.Get("clickhouse")
				if clickhouse != "" && clickhouse != "-" {
					s := strings.Split(clickhouse, ";")
					newKey := field
					for _, key := range s {
						r := strings.Split(key, ":")

						if len(r) > 1 {

							n := r[0]
							v := r[1]

							if n == "type" {
								newKey += " " + v
							} else if n == "default" {
								newKey += " DEFAULT " + v
							}
						} else {

							if key == "key" {
								masterKey = field
							} else if key == "order" {
								orders = append(orders, field)
							} else {
								newKey += key
							}
						}
					}

					fields = append(fields, newKey)
				}
			}
		}

		sql += strings.Join(fields[:], ",\n")
		sql += ")\n engine = " + engine

		if masterKey != "" {
			sql += "(" + masterKey + ")"
		}

		if len(orders) > 0 {
			sql += " ORDER BY (" + strings.Join(orders[:], ",") + ")"
		}

		sql += ";"
		return sql
	}

	panic(fmt.Errorf("DBFields requires a struct or a map, found: %s", v.Kind().String()))
}
