package apirouterv1

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	controllerv1 "github.com/qxip/cloki-go/controller/v1"
	"github.com/qxip/cloki-go/data/service"
	"github.com/qxip/cloki-go/model"
)

func RoutePushApis(acc *echo.Group, dataSession *sqlx.DB, goCache *cache.Cache, dbTime, bufferSize, dbBulk int) {
	// initialize service of user
	pusService := service.PushService{ServiceData: service.ServiceData{Session: dataSession},
		GoCache: goCache,
		TSCh:    make(chan *model.TableTimeSeries, bufferSize),
		SPCh:    make(chan *model.TableSample, bufferSize),
		DBTimer: time.Duration(dbTime) * time.Second,
		DBBulk:  dbBulk,
	}
	// initialize user controller
	urc := controllerv1.PushController{
		PushService: &pusService,
	}
	// user login
	go pusService.Insert()
	pusService.ReloadFingerprints()
	acc.POST("/push", urc.PushStream)

}
