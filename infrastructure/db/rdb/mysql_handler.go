package rdb

import (
	"boilerplate/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MysqlHandler struct {
	Conn *gorm.DB
}

var mysqlInstance = &MysqlHandler{}

func Init(config config.Mysql) error {
	fmt.Println("aa:",config)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return err
	}

	mysqlInstance.Conn = db
	return nil
}

func GetInstance() *gorm.DB {
	return mysqlInstance.Conn
}
