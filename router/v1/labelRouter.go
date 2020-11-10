package apirouterv1

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	controllerv1 "gitlab.com/qxip/cloki/controller/v1"
	"gitlab.com/qxip/cloki/data/service"
)

func RouteLabelApis(acc *echo.Group, dataSession *sqlx.DB, goCache *cache.Cache) {
	// initialize service of user
	labelService := service.LabelService{ServiceData: service.ServiceData{Session: dataSession},
		GoCache: goCache,
	}
	// initialize user controller
	lr := controllerv1.LabelController{
		LabelService: &labelService,
	}
	// user login
	acc.GET("/label", lr.GetLabels)
	acc.GET("/label/:name/values", lr.LabelValsByKey)

}
