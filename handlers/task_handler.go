package handlers

import (
	"net/http"
	"strconv"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	tasks, err := models.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, tasks)
}

func CreateTask(ctx *gin.Context) {
	var reqTask models.ReqTask
	if err := ctx.BindJSON(&reqTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	task, err := reqTask.CreateTask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, task)
}

func UpdateTask(ctx *gin.Context) {
	var reqTask models.ReqUpdateTask
	if err := ctx.BindJSON(&reqTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := reqTask.UpdateTask(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.String(http.StatusAccepted, "OK")
}

func DeleteTask(ctx *gin.Context) {
	var id int
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := models.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.String(http.StatusAccepted, "OK")
}
