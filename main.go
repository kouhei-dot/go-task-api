package main

import (
	"fmt"
	"net/http"
	"todo-api/config"
	"todo-api/controller"
)

func init() {
	config.LoadConfig()
}

func main() {
	controller.Router()
	// log.Fatal(http.ListenAndServe(":8080", nil))
	err := http.ListenAndServe(":"+config.Conf.Web.Port, nil)
	if err != nil {
		fmt.Println("start server")
	} else {
		fmt.Println("failed")
	}
}
