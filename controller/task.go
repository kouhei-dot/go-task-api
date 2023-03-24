package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todo-api/models"
)

func Handler(writer http.ResponseWriter, req *http.Request) {
	var errHandle = func(statusCode int, err error) {
		writer.WriteHeader(statusCode)
		writer.Write([]byte("Failed"))
		log.Println(err)
	}

	log.Println("Handler start")
	writer.Header().Add("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case http.MethodGet:
		var tasks, dbErr = models.GetTasks()
		if dbErr != nil {
			errHandle(http.StatusNotFound, dbErr)
			return
		}

		var res, jsonErr = json.Marshal(tasks)
		if jsonErr != nil {
			errHandle(http.StatusInternalServerError, jsonErr)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(res)

	case http.MethodPost:
		var reqTask models.ReqTask
		if err := json.NewDecoder(req.Body).Decode(&reqTask); err != nil {
			errHandle(http.StatusInternalServerError, err)
			return
		}

		task, err := reqTask.CreateTask()
		if err != nil {
			errHandle(http.StatusInternalServerError, err)
			return
		}

		res, err := json.Marshal(task)
		if err != nil {
			errHandle(http.StatusInternalServerError, err)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		writer.Write(res)

	case http.MethodPatch:
		var task models.ReqUpdateTask
		if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
			errHandle(http.StatusInternalServerError, err)
			return
		}

		if err := task.UpdateTask(); err != nil {
			errHandle(http.StatusInternalServerError, err)
			return
		}

		writer.WriteHeader(http.StatusAccepted)

	case http.MethodDelete:
		id, err := strconv.Atoi(strings.SplitAfter(req.URL.Path, "/")[1])
		if err != nil {
			errHandle(http.StatusInternalServerError, err)
			return
		}
		models.DeleteTask(id)
		writer.WriteHeader(http.StatusAccepted)

	case http.MethodOptions:
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Add("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))

	default:
		writer.WriteHeader(http.StatusBadRequest)
	}
}
