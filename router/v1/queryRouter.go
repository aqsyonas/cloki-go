package apirouterv1

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	controllerv1 "github.com/qxip/cloki-go/controller/v1"
	"github.com/qxip/cloki-go/data/service"
)

func RouteQueryApis(acc *echo.Group, dataSession *sqlx.DB) {
	// initialize service of user
	queryService := service.QueryService{ServiceData: service.ServiceData{Session: dataSession}}
	// initialize user controller
	qc := controllerv1.QueryController{
		QueryService: &queryService,
	}
	// user login
	acc.GET("/query", qc.GetQuery)
	acc.GET("/query_range", qc.QueryRange)

}
