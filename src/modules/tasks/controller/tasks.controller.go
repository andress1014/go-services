package taskController

import (
	"net/http"

	"github.com/andress1014/go-gorm-crud/src/config/handler"
	domainTask "github.com/andress1014/go-gorm-crud/src/modules/tasks/domain"
	"github.com/gorilla/mux"
)

func GetTaskController(w http.ResponseWriter, r *http.Request) {
	result := domainTask.GetTasksDomain()

	handler.HandleResponse(w, http.StatusOK, result)
}

func GetOneTasksController(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["taskId"]
	result := domainTask.GetOneTasksDomain(w, taskId)

	if result.Error != nil {
		handler.HandleError(w, http.StatusNotFound, result.Error.Error())
		return
	}

	handler.HandleResponse(w, http.StatusOK, result.Task)
}
