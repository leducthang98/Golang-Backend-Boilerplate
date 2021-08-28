package main

import (
	"boilerplate/config"
	"boilerplate/server"
	"context"
)

func main() {
	srvCfg := new(config.Server)

	srvCfg.Mysql = new(config.Mysql)
	// env load
	srvCfg.Port = 3000
	srvCfg.Mysql.Host = "localhost"
	srvCfg.Mysql.Port = 3306
	srvCfg.Mysql.User = "root"
	srvCfg.Mysql.Password = "password"
	srvCfg.Mysql.DatabaseName = "default"

	srv := server.New(srvCfg)

	if err := srv.Start(context.Background()); err != nil {
		panic(err)
	}

	defer srv.Close()
}
