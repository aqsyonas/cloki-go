package apirouterv1

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	controllerv1 "gitlab.com/qxip/cloki/controller/v1"
	"gitlab.com/qxip/cloki/data/service"
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
