package controller

import (
	"boilerplate/pkg/api_utils"
	"boilerplate/pkg/constant"
	"boilerplate/server/container"
	"boilerplate/server/middleware"
	"boilerplate/service/user"
	"github.com/labstack/echo/v4"
)

const (
	userRoute = "/user"
)

type UserController struct {
	userService *user.UserService
}

func MountUserController(e *echo.Echo) {
	userController := UserController{
		userService: container.GetInstance().Get(constant.UserIocService).(*user.UserService),
	}

	g := e.Group(userRoute)

	g.GET("/all", userController.getAllUser, middleware.DefaultMiddleware("jwt secret"))
}

func (u *UserController) getAllUser(c echo.Context) error {
	listUserResources, err := u.userService.GetAll()
	if err != nil {
		return api_utils.ErrBadRequest(c, 1, err.Error())
	}
	return api_utils.Ok(c, listUserResources)
}
