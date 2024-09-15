package main

import (
	"net/http"

	"github.com/andress1014/go-gorm-crud/src/db"
	models "github.com/andress1014/go-gorm-crud/src/model"
	"github.com/andress1014/go-gorm-crud/src/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", jsonResponse(routes.GetUsersHandler)).Methods("GET")
	r.HandleFunc("/users/{id}", jsonResponse(routes.GetUserHandler)).Methods("GET")
	r.HandleFunc("/users", jsonResponse(routes.PostUsersHandler)).Methods("POST")
	r.HandleFunc("/users/{id}", jsonResponse(routes.DeleteUsersHandler)).Methods("DELETE")
	//Task routes
	r.HandleFunc("/tasks", jsonResponse(routes.GetTasksHandler)).Methods("GET")
	r.HandleFunc("/tasks/{id}", jsonResponse(routes.GetTaskHandler)).Methods("GET")
	r.HandleFunc("/tasks", jsonResponse(routes.CreateTaskHandler)).Methods("POST")
	r.HandleFunc("/tasks/{id}", jsonResponse(routes.DeleteTaskHandler)).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}

func jsonResponse(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}
