package routes

import (
	"encoding/json"
	"net/http"

	"github.com/andress1014/go-gorm-crud/src/config/handler"
	"github.com/andress1014/go-gorm-crud/src/db"
	models "github.com/andress1014/go-gorm-crud/src/model"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)

	handler.HandleResponse(w, http.StatusOK, tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		handler.HandleError(w, http.StatusBadRequest, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.DB.Create(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.HandleResponse(w, http.StatusCreated, tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	params := mux.Vars(r)
	db.DB.First(&tasks, params["id"])

	if tasks.ID == 0 {
		handler.HandleError(w, http.StatusNotFound, "Record not found")
		return
	}
	handler.HandleResponse(w, http.StatusOK, tasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	params := mux.Vars(r)
	db.DB.First(&tasks, params["id"])

	if tasks.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "record not found"}`))
		return
	}

	db.DB.Unscoped().Delete(&tasks)
	handler.HandleResponse(w, http.StatusOK, tasks)
}
