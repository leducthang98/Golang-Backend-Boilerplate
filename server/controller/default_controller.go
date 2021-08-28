package controller

import "github.com/labstack/echo/v4"

const (
	defaultRoute = "/"
)

type DefaultController struct {
}

func MountDefaultController(e *echo.Echo) {
	defaultController := DefaultController{

	}

	g := e.Group(defaultRoute)

	g.GET("ping", defaultController.ping)
}

func (d *DefaultController) ping(e echo.Context) error {
	e.JSON(200, "pong")
	return nil
}
