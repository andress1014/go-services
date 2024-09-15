package routes

import (
	"encoding/json"
	"net/http"

	"github.com/andress1014/go-gorm-crud/src/db"
	models "github.com/andress1014/go-gorm-crud/src/model"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.DB.Create(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	params := mux.Vars(r)
	db.DB.First(&tasks, params["id"])

	if tasks.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "record not found"}`))
	}
	json.NewEncoder(w).Encode(tasks)

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
	w.WriteHeader(http.StatusOK)
}
