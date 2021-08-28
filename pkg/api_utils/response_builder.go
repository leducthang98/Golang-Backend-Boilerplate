package api_utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ResponseForm struct {
	// code, message, data
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	ErrorCode    int    `json:"error_code""`
	ErrorMessage string `json:"error_message""`
}

func Ok(c echo.Context, data interface{}) error {
	c.JSON(http.StatusOK, ResponseForm{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
	return nil
}

func ErrBadRequest(c echo.Context, errorCode int, errorMessage string) error {
	c.JSON(http.StatusBadRequest, Error{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	})
	return nil
}
