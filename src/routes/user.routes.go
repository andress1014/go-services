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
	handler.HandleResponse(w, http.StatusOK, users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var users models.User
	params := mux.Vars(r)
	db.DB.First(&users, params["id"])

	if users.ID == 0 {
		handler.HandleError(w, http.StatusNotFound, "record not found")
		return
	}
	handler.HandleResponse(w, http.StatusOK, users)
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		handler.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.HandleResponse(w, http.StatusCreated, user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users models.User
	params := mux.Vars(r)
	db.DB.First(&users, params["id"])

	if users.ID == 0 {
		handler.HandleError(w, http.StatusNotFound, "record not found")
		return
	}

	db.DB.Unscoped().Delete(&users)
	w.WriteHeader(http.StatusOK)
}
