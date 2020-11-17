package controllerv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qxip/cloki-go/data/service"
	httpresponse "github.com/qxip/cloki-go/network/response"
)

type SeriesController struct {
	Controller
	SeriesService *service.SeriesService
}

func (sc *SeriesController) GetSeries(c echo.Context) error {
	reply := sc.SeriesService.GetSeries()
	return httpresponse.CreateBadResponseWithJson(&c, http.StatusOK, []byte(reply))
}
