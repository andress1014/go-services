package domainTask

import (
	"errors"
	"net/http"

	"github.com/andress1014/go-gorm-crud/src/db"
	models "github.com/andress1014/go-gorm-crud/src/model"
)

func GetTasksDomain() []models.Task {
	var tasks []models.Task
	db.DB.Find(&tasks)

	return tasks
}

type TaskResult struct {
	Task  models.Task
	Error error
}

func GetOneTasksDomain(w http.ResponseWriter, id string) TaskResult {
	var task models.Task
	db.DB.First(&task, id)
	if task.ID == 0 {
		return TaskResult{Error: errors.New("record not found")}
	}
	return TaskResult{Task: task}
}
