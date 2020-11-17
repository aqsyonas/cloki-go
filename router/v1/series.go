package apirouterv1

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	controllerv1 "github.com/qxip/cloki-go/controller/v1"
	"github.com/qxip/cloki-go/data/service"
)

func RouteSeriesApis(acc *echo.Group, dataSession *sqlx.DB) {
	// initialize service of user
	seriesService := service.SeriesService{ServiceData: service.ServiceData{Session: dataSession}}
	// initialize user controller
	sc := controllerv1.SeriesController{
		SeriesService: &seriesService,
	}
	// user login
	acc.GET("/series", sc.GetSeries)

}
