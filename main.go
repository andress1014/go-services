package main

import (
	"log"
	"net/http"

	"github.com/andress1014/go-gorm-crud/src/db"
	models "github.com/andress1014/go-gorm-crud/src/model"
	taskRouter "github.com/andress1014/go-gorm-crud/src/modules/tasks/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Conectar la base de datos
	db.DBConnection()

	// Migrar los modelos
	db.DB.AutoMigrate(&models.Task{})
	db.DB.AutoMigrate(&models.User{})

	// Crear un nuevo enrutador
	r := mux.NewRouter()

	// Registrar las rutas de tareas bajo el prefijo /task
	taskRouter.SetupTaskRoutes(r)

	// Iniciar el servidor
	log.Println("Server running on http://localhost:8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
