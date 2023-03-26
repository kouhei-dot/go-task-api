package main

import (
	"net/http"
	"todo-api/config"
	"todo-api/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
}

func main() {
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		if ctx.Request.Method == http.MethodOptions {
			ctx.String(http.StatusOK, "OK")
		}
	})
	router.GET("/tasks", handlers.GetTasks)
	router.POST("/task", handlers.CreateTask)
	router.PATCH("/task", handlers.UpdateTask)
	router.DELETE("/task/:id", handlers.DeleteTask)
	router.Run(":" + config.Conf.Web.Port)
}
