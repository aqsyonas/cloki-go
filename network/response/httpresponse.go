package httpresponse

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CreateBadResponse(c *echo.Context, requestCode int, message string) error {
	localC := *c
	response := fmt.Sprintf("{\"data\":{},\"message\":%q}", message)
	return localC.JSONBlob(requestCode, []byte(response))
}

func CreateSuccessResponse(c *echo.Context, requestCode int, message string) error {
	localC := *c
	return localC.JSONBlob(requestCode, []byte(message))
}

func CreateBadResponseWithJson(c *echo.Context, requestCode int, message []byte) error {
	localC := *c
	return localC.JSONBlob(requestCode, message)
}

func CreateSuccessResponseWithJson(c *echo.Context, requestCode int, message []byte) error {
	localC := *c
	return localC.JSONBlob(requestCode, []byte(message))
}

func CreateSuccessResponseWTBody(c *echo.Context, requestCode int) error {
	localC := *c
	return localC.NoContent(requestCode)
}

func CreateSuccessResponseWithCSV(c *echo.Context, requestCode int, filename string, body []byte) error {
	localC := *c

	res := localC.Response()
	header := res.Header()
	header.Set(echo.HeaderContentType, echo.MIMEOctetStream)
	header.Set(echo.HeaderContentDisposition, "attachment; filename="+filename)
	header.Set("Content-Transfer-Encoding", "binary")
	header.Set("Expires", "0")

	return localC.Blob(requestCode, "text/csv", body)

	/* res.WriteHeader(http.StatusOK)
	res.Write(body)
	res.Flush()
	return nil
	*/
	/*
		res := localC.Response()
		header := res.Header()
		header.Set(echo.HeaderContentDisposition, "attachment; filename="+filename)
		header.Set("Expires", "0")
		return localC.Blob(requestCode, "text/csv", body)
	*/
}
