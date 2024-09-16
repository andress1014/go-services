package taskRouter

import (
	taskController "github.com/andress1014/go-gorm-crud/src/modules/tasks/controller"
	"github.com/gorilla/mux"
)

func SetupTaskRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", taskController.GetTaskController).Methods("GET")
	r.HandleFunc("/tasks/{taskId}", taskController.GetOneTasksController).Methods("GET")
	// r.HandleFunc("/tasks", jsonResponse(routes.CreateTaskHandler)).Methods("POST")
	// r.HandleFunc("/tasks/{id}", jsonResponse(routes.DeleteTaskHandler)).Methods("DELETE")
}
