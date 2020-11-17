package apirouterv1

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	controllerv1 "github.com/qxip/cloki-go/controller/v1"
	"github.com/qxip/cloki-go/data/service"
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
