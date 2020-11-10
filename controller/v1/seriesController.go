package controllerv1

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/qxip/cloki/data/service"
	httpresponse "gitlab.com/qxip/cloki/network/response"
	"net/http"
)

type SeriesController struct {
	Controller
	SeriesService *service.SeriesService
}

func (sc *SeriesController) GetSeries(c echo.Context) error {
	reply := sc.SeriesService.GetSeries()
	return httpresponse.CreateBadResponseWithJson(&c, http.StatusOK, []byte(reply))
}
