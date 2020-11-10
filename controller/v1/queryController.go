package controllerv1

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/qxip/cloki/data/service"
	httpresponse "gitlab.com/qxip/cloki/network/response"
	"net/http"
)

type QueryController struct {
	Controller
	QueryService *service.QueryService
}

func (qc *QueryController) GetQuery(c echo.Context) error {
	reply := qc.QueryService.GetQuery()
	return httpresponse.CreateBadResponseWithJson(&c, http.StatusOK, []byte(reply))
}

func (qc *QueryController) QueryRange(c echo.Context) error {
	query := c.QueryParam("query")
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	reply := qc.QueryService.QueryRange(query, start, end)
	return httpresponse.CreateBadResponseWithJson(&c, http.StatusOK, []byte(reply))
}
