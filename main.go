// CLoki-App
//
// CLoki-App User interface for WEB AI
//
//     Schemes: http
//     Host: localhost:9080
//     BasePath: /api/v3
//     Version: 1.1.2
//     License: AGPL https://www.gnu.org/licenses/agpl-3.0.en.html
//	   Copyright: QXIP B.V. 2019-2020
//     Contact: Aqs <aqsyounas@gmail.com>
//     Contact: Alexandr Dubovikov <alexandr.dubovikov@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/qxip/cloki-go/migration"
	"github.com/qxip/cloki-go/migration/jsonschema"
	"github.com/qxip/cloki-go/utils/heputils"
	"github.com/qxip/cloki-go/utils/logger"

	"os"

	"github.com/ClickHouse/clickhouse-go"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apirouterv1 "github.com/qxip/cloki-go/router/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

//CustomValidator function
type CustomValidator struct {
	validator *validator.Validate
}

// validate function
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var appFlags CommandLineFlags
var authType string

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

//params for Flags
type CommandLineFlags struct {
	InitializeDB          *bool   `json:"initialize_db"`
	CreateClokiDB         *bool   `json:"create_cloki_db"`
	CreateTableCLokiDB    *bool   `json:"create_table_cloki"`
	CreateCLokiUser       *bool   `json:"create_cloki_user"`
	DeleteCLokiUser       *bool   `json:"delete_cloki_user"`
	ShowVersion           *bool   `json:"version"`
	ShowHelpMessage       *bool   `json:"help"`
	DatabaseRootUser      *string `json:"root_user"`
	DatabaseRootPassword  *string `json:"root_password"`
	DatabaseHost          *string `json:"root_host"`
	DatabasePort          *int    `json:"root_port"`
	DatabaseRootDB        *string `json:"root_db"`
	DatabaseCLokiNode     *string `json:"cloki_node"`
	DatabaseCLokiUser     *string `json:"cloki_user"`
	DatabaseCLokiPassword *string `json:"cloki_password"`
	DatabaseCloki         *string `json:"db_cloki"`
	PathCLokiappConfig    *string `json:"path_clokiapp"`
	LogPathCLokiapp       *string `json:"path_log_clokiapp"`
	LogName               *string `json:"log_name_clokiapp"`
	APIPrefix             *string `json:"api_prefix"`
}

//params for  Services
type ServicesObject struct {
	dataDBSession *sqlx.DB
}

var servicesObject ServicesObject

/* init flags */
func initFlags() {
	appFlags.InitializeDB = flag.Bool("initialize_db", false, "initialize the database and create all tables")
	appFlags.CreateClokiDB = flag.Bool("create-cloki-db", false, "create cloki db")
	appFlags.CreateTableCLokiDB = flag.Bool("create-table-db-cloki", false, "create table in db config")

	appFlags.CreateCLokiUser = flag.Bool("create-cloki-user", false, "create cloki user")
	appFlags.DeleteCLokiUser = flag.Bool("delete-cloki-user", false, "delete cloki user")
	appFlags.ShowVersion = flag.Bool("version", false, "show version")

	appFlags.ShowHelpMessage = flag.Bool("help", false, "show help")
	appFlags.DatabaseRootUser = flag.String("database-root-user", "default", "database-root-user")
	appFlags.DatabaseRootPassword = flag.String("database-root-password", "clickPass", "database-root-password")
	appFlags.DatabaseHost = flag.String("database-host", "127.0.0.1", "database-host")
	appFlags.DatabasePort = flag.Int("database-port", 9000, "database-port")
	appFlags.DatabaseRootDB = flag.String("database-root-db", "default", "database-root-db")
	appFlags.DatabaseCLokiNode = flag.String("database-cloki-node", "localnode", "database-cloki-node")
	appFlags.DatabaseCLokiUser = flag.String("database-cloki-user", "cloki_user", "database-cloki-user")
	appFlags.DatabaseCLokiPassword = flag.String("database-cloki-password", "cloki_password", "database-cloki-password")
	appFlags.DatabaseCloki = flag.String("database-cloki-data", "cloki", "database-cloki")

	appFlags.PathCLokiappConfig = flag.String("clokiapp-config-path", "/usr/local/cloki/etc", "the path to the clokiapp config file")
	appFlags.LogName = flag.String("clokiapp-log-name", "", "the name prefix of the log file.")
	appFlags.LogPathCLokiapp = flag.String("clokiapp-log-path", "", "the path for the log file.")
	appFlags.APIPrefix = flag.String("clokiapp-api-prefix", "", "API prefix.")

	flag.Parse()
}

func main() {

	//init flags
	initFlags()

	/* first check admin flags */
	checkHelpVersionFlags()

	// read system configurations and expose through viper
	readConfig()

	logPath := viper.GetString("system_settings.logpath")
	logName := viper.GetString("system_settings.logname")
	logLevel := viper.GetString("system_settings.loglevel")
	logStdout := viper.GetBool("system_settings.logstdout")

	if *appFlags.LogPathCLokiapp != "" {
		logPath = *appFlags.LogPathCLokiapp
	} else if logPath == "" {
		logPath = "log"
	}

	if *appFlags.LogName != "" {
		logName = *appFlags.LogName
	} else if logPath == "" {
		logName = "cloki.log"
	}

	// initialize logger
	logger.InitLogger(logPath, logName, logLevel, logStdout)

	/* first check admin flags */
	checkAdminFlags()

	servicesObject.dataDBSession = getDataDBSession()

	if *appFlags.CreateTableCLokiDB {
		nameCLokiConfig := viper.GetString("database_config.name")
		migration.CreateClokiTables(servicesObject.dataDBSession, nameCLokiConfig)
		os.Exit(0)
	}

	defer servicesObject.dataDBSession.Close()

	if versionPg, err := migration.CheckVersion(servicesObject.dataDBSession); err != nil {
		heputils.Colorize(heputils.ColorRed, "\r\nVersion of DB couldn't be retrieved\r\n")
	} else if (versionPg / 10000) < jsonschema.MinimumClickHouse {
		heputils.Colorize(heputils.ColorRed, fmt.Sprintf("\r\nYou don't have required version of Clickhouse. Please install minimum: %d\r\n", jsonschema.MinimumClickHouse))
	} else {
		heputils.Colorize(heputils.ColorBlue, fmt.Sprintf("\r\nClickhouse version: %d.%d\r\n", versionPg/10000, versionPg%10000))
	}

	// configure to serve WebServices
	configureAsHTTPServer()

}

func configureAsHTTPServer() {

	e := echo.New()
	// add validation
	e.Validator = &CustomValidator{validator: validator.New()}
	// Middleware
	if httpDebugEnable := viper.GetBool("http_settings.debug"); httpDebugEnable {
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{Skipper: skipper, Handler: bodyDumpHandler}))
	}

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	/* hide bunner */
	e.HideBanner = true

	// perform routing for v1 version of web apis
	performV1APIRouting(e)

	heputils.Colorize(heputils.ColorGreen, fmt.Sprintf("Version: %s %s", getName(), getVersion()))

	if viper.GetBool("https_settings.enable") {
		httpsHost := viper.GetString("https_settings.host")
		httpsPort := viper.GetString("https_settings.port")
		httpsURL := fmt.Sprintf("%s:%s", httpsHost, httpsPort)

		httpsCert := viper.GetString("https_settings.cert")
		httpsKey := viper.GetString("https_settings.key")
		//Doc Swagger for future. For now - external
		/* e.GET("/swagger/*", echoSwagger.WrapHandler)
		 */
		e.Logger.Fatal(e.StartTLS(httpsURL, httpsCert, httpsKey))
	} else {
		httpHost := viper.GetString("http_settings.host")
		httpPort := viper.GetString("http_settings.port")
		httpURL := fmt.Sprintf("%s:%s", httpHost, httpPort)

		//Doc Swagger for future. For now - external
		/* e.GET("/swagger/*", echoSwagger.WrapHandler)
		 */
		e.Logger.Fatal(e.Start(httpURL))
	}
}

func performV1APIRouting(e *echo.Echo) {

	// accessible web services will fall in this group

	goCache := cache.New(30*time.Minute, 10*time.Minute)

	acc := e.Group("/loki/api/v1")

	dbTime := viper.GetInt("database_data.dbTime")
	bufferSize := viper.GetInt("database_data.bufferSize")
	dbBulk := viper.GetInt("database_data.dbBulk")

	apirouterv1.RoutePushApis(acc, servicesObject.dataDBSession, goCache, dbTime, bufferSize, dbBulk)
	apirouterv1.RouteLabelApis(acc, servicesObject.dataDBSession, goCache)
	apirouterv1.RouteQueryApis(acc, servicesObject.dataDBSession)

}

// getSession creates a new postgres session and panics if connection error occurs
func getDataDBSession() *sqlx.DB {
	user := viper.GetString("database_data.user")
	password := viper.GetString("database_data.pass")
	name := viper.GetString("database_data.name")
	host := viper.GetString("database_data.host")
	port := viper.GetInt("database_data.port")
	read_timeout := viper.GetInt("database_data.read_timeout")
	write_timeout := viper.GetInt("database_data.write_timeout")

	if read_timeout < 10 {
		read_timeout = 30
	}

	if write_timeout < 30 {
		write_timeout = 30
	}

	if port == 0 {
		port = 9000
	}

	connectString := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s&read_timeout=%d&write_timeout=%d&compress=true&debug=true",
		host, port, user, password, name, read_timeout, write_timeout)

	logrus.Info(fmt.Sprintf("Connecting to the config: [%s, %s, %s, %d]\n", host, user, name, port))
	db, err := sqlx.Open("clickhouse", connectString)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			logrus.Error(fmt.Sprintf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace))
		} else {
			logrus.Info("ping db config", err)
		}
		return nil
	}

	logrus.Info("----------------------------------- ")
	logrus.Info("*** Database Config Session created *** ")
	logrus.Info("----------------------------------- ")

	return db
}

func readConfig() {

	// Getting constant values
	if configEnv := os.Getenv("CLOKIAPPENV"); configEnv != "" {
		viper.SetConfigName("clokiapp_config_" + configEnv)
	} else {
		viper.SetConfigName("clokiapp_config")
	}
	viper.SetConfigType("json")

	if configPath := os.Getenv("CLOKIAPPPATH"); configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath(*appFlags.PathCLokiappConfig)
	}
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No configuration file loaded: ", err)
		logrus.Errorln("No configuration file loaded - using defaults")
		panic("DB configuration file not found: ")
	}
}

// middle ware handler
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {

	logrus.Println("================================")
	logrus.Println("--------request body-------")
	printBody(reqBody)
	logrus.Println("---------------------------")

	logrus.Println("-------- response body --------")
	printBody(resBody)
	logrus.Println("-------------------------------")
	logrus.Println("=================================")
}

// body dump skipper
func skipper(c echo.Context) bool {

	if c.Request().Method == "POST" {
		logrus.Println(c.Request().URL.Path)
		return true
	}

	return false
}

// private method
func printBody(obj []byte) {

	logrus.WithFields(logrus.Fields{
		"json": string(obj),
	}).Info("Payload")
}

func printRequest(request *http.Request) {

	logrus.WithFields(logrus.Fields{
		"HOST":       request.Host,
		"PATH":       request.URL.Path,
		"METHOD":     request.Method,
		"QueryParam": request.URL.Query().Encode(),
	}).Info("Request")
}

func checkHelpVersionFlags() {

	if *appFlags.ShowHelpMessage {
		flag.Usage()
		os.Exit(0)
	}

	if *appFlags.ShowVersion {
		fmt.Printf("VERSION: %s\r\n", getVersion())
		os.Exit(0)
	}
}

func checkAdminFlags() {

	if *appFlags.InitializeDB {
		initDB()
	}

	/* start creating pgsql user */
	if *appFlags.CreateCLokiUser {

		rootDb, err := migration.GetDataRootDBSession(appFlags.DatabaseRootUser,
			appFlags.DatabaseRootPassword,
			appFlags.DatabaseRootDB,
			appFlags.DatabaseHost,
			appFlags.DatabasePort)

		if err != nil {
			logrus.Error("Couldn't establish connection. Please be sure you can have correct password", err)
			logrus.Error("Try run: sudo -u postgres psql -c \"ALTER USER postgres PASSWORD 'postgres';\"")
			panic(err)
		}

		defer rootDb.Close()

		migration.CreateNewUser(rootDb, appFlags.DatabaseCLokiUser, appFlags.DatabaseCLokiPassword)

		os.Exit(0)
		/* start drop pgsql user */
	} else if *appFlags.CreateClokiDB {

		rootDb, err := migration.GetDataRootDBSession(appFlags.DatabaseRootUser,
			appFlags.DatabaseRootPassword,
			appFlags.DatabaseRootDB,
			appFlags.DatabaseHost,
			appFlags.DatabasePort)

		if err != nil {
			logrus.Error("Couldn't establish connection. Please be sure you can have correct password", err)
			logrus.Error("Try run: sudo -u postgres psql -c \"ALTER USER postgres PASSWORD 'postgres';\"")
			panic(err)
		}

		defer rootDb.Close()
		migration.CreateClokiDB(rootDb, appFlags.DatabaseCloki, appFlags.DatabaseCLokiUser)
		os.Exit(0)
		/* start creating pgsql user */
	}
}

func initDB() {

	rootDb, err := migration.GetDataRootDBSession(
		appFlags.DatabaseRootUser,
		appFlags.DatabaseRootPassword,
		appFlags.DatabaseRootDB,
		appFlags.DatabaseHost,
		appFlags.DatabasePort,
	)

	if err != nil {
		logrus.Error("Couldn't establish connection. Please be sure you can have correct password", err)
		logrus.Error("Try run: sudo -u postgres psql -c \"ALTER USER postgres PASSWORD 'postgres';\"")
		panic(err)
	}

	defer rootDb.Close()

	migration.CreateNewUser(rootDb, appFlags.DatabaseCLokiUser, appFlags.DatabaseCLokiPassword)

	migration.CreateClokiDB(rootDb, appFlags.DatabaseCloki, appFlags.DatabaseCLokiUser)

	databaseDb, err := migration.GetDataRootDBSession(
		appFlags.DatabaseRootUser,
		appFlags.DatabaseRootPassword,
		appFlags.DatabaseCloki,
		appFlags.DatabaseHost,
		appFlags.DatabasePort,
	)

	if err != nil {
		logrus.Error("Couldn't establish connection to databaseDb. Please be sure you can have correct password", err)
		logrus.Error("Try run: sudo -u postgres psql -c \"ALTER USER postgres PASSWORD 'postgres';\"")
		panic(err)
	}

	defer databaseDb.Close()

	servicesObject.dataDBSession = getDataDBSession()
	defer servicesObject.dataDBSession.Close()
	nameHepicConfig := viper.GetString("database_config.name")

	migration.CreateClokiTables(servicesObject.dataDBSession, nameHepicConfig)

	os.Exit(0)
}
