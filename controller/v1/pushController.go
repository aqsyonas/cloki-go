package controllerv1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/qxip/cloki-go/data/service"
	"github.com/qxip/cloki-go/model"
	httpresponse "github.com/qxip/cloki-go/network/response"
	"github.com/qxip/cloki-go/system/webmessages"
)

type PushController struct {
	Controller
	PushService *service.PushService
}

func (pc *PushController) PushStream(c echo.Context) error {
	var req model.PushRequest
	if err := c.Bind(&req); err != nil {
		logrus.Error(err.Error())
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, webmessages.UserRequestFormatIncorrect)
	}
	//validate input request body
	if err := c.Validate(req); err != nil {
		logrus.Error(err.Error())
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, err.Error())
	}
	if err := pc.PushService.PushStream(req); err != nil {
		return httpresponse.CreateBadResponseWithJson(&c, http.StatusUnauthorized, []byte(err.Error()))
	}

	return httpresponse.CreateSuccessResponseWTBody(&c, http.StatusNoContent)
}
