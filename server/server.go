package server

import (
	"boilerplate/config"
	"boilerplate/infrastructure/db/rdb"
	"boilerplate/infrastructure/repository/user"
	"boilerplate/pkg/constant"
	"boilerplate/server/container"
	"boilerplate/server/controller"
	user2 "boilerplate/service/user"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"log"
	"strconv"
)

const (
	appName = "echo-boilerplate"
)

type Server struct {
	cfg *config.Server
}

func New(cfg *config.Server) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start(ctx context.Context) (err error) {

	// init echo
	e := echo.New()

	// init db connection
	rdb.Init(*s.cfg.Mysql)

	// ioc
	instances := []di.Def{} // list instance to inject
	repositories := s.MountRepository()
	services := s.MountService()
	instances = append(instances, repositories...)
	instances = append(instances, services...)
	err = container.Init(instances...) // co full instance repo service

	// mount routes
	s.MountController(e)

	// start
	port := strconv.Itoa(int(s.cfg.Port))
	e.Logger.Fatal(e.Start(":" + port))

	return
}

func (s *Server) MountRepository() []di.Def {
	return []di.Def{
		{
			Name:     constant.UserIocRepository,
			Unshared: false,
			Build: func(ctn di.Container) (interface{}, error) {
				return user.NewUserRepository(), nil
			},
		},
	}
}

func (s *Server) MountService() []di.Def {
	return []di.Def{
		{
			Name:     constant.UserIocService,
			Unshared: false,
			Build: func(ctn di.Container) (interface{}, error) {
				return user2.NewUserService(), nil
			},
		},
	}
}

func (s *Server) MountController(e *echo.Echo) {
	controller.MountDefaultController(e)
	controller.MountUserController(e)
}

func (s *Server) Close() {
	log.Print("Server closed")
}
