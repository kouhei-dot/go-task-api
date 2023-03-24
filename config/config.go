package config

import (
	"fmt"
	"log"
	"todo-api/models"

	"gopkg.in/ini.v1"
)

type Config struct {
	Web Web
	Db  Db
}

type Web struct {
	Port string
}

type Db struct {
	User     string
	Port     string
	DbName   string
	Password string
	Host     string
}

var Conf Config

func LoadConfig() {
	conf, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	Conf.Web.Port = conf.Section("web").Key("port").MustString("8080")
	db := conf.Section("db")
	Conf.Db.User = db.Key("user").String()
	Conf.Db.Port = db.Key("port").String()
	Conf.Db.DbName = db.Key("db").String()
	Conf.Db.Password = db.Key("password").String()
	Conf.Db.Host = db.Key("host").MustString("127.0.0.1")

	models.ConnectionString = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		Conf.Db.User, Conf.Db.Password, Conf.Db.Host, Conf.Db.Port, Conf.Db.DbName,
	)
}
