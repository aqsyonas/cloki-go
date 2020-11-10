package controllerv1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gitlab.com/qxip/cloki/data/service"

	httpresponse "gitlab.com/qxip/cloki/network/response"
)

type LabelController struct {
	Controller
	LabelService *service.LabelService
}

func (lb *LabelController) GetLabels(c echo.Context) error {

	label := lb.LabelService.GetLabels()
	return httpresponse.CreateSuccessResponseWithJson(&c, http.StatusCreated, []byte(label))
}

func (lb *LabelController) LabelValsByKey(c echo.Context) error {

	label := lb.LabelService.LabelValsByKey(c.Param("name"))
	return httpresponse.CreateSuccessResponseWithJson(&c, http.StatusCreated, []byte(label))
}
