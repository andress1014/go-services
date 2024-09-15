package routes

import (
	"encoding/json"
	"net/http"

	"github.com/andress1014/go-gorm-crud/src/config/handler"
	"github.com/andress1014/go-gorm-crud/src/db"
	models "github.com/andress1014/go-gorm-crud/src/model"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var users models.User
	params := mux.Vars(r)
	db.DB.First(&users, params["id"])

	if users.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "record not found"}`))
		return
	}
	json.NewEncoder(w).Encode(users)
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.HandleResponse(w, http.StatusCreated, user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users models.User
	params := mux.Vars(r)
	db.DB.First(&users, params["id"])

	if users.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "record not found"}`))
		return
	}

	db.DB.Unscoped().Delete(&users)
	w.WriteHeader(http.StatusOK)
}
