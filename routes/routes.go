package routes

import (
	"GoAssignment2/db"
	"GoAssignment2/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes(dbConn *db.Database, gormDBConn *db.AdvancedDatabase) *mux.Router {
	r := mux.NewRouter()

	userHandler := &handlers.UserHandler{DB: dbConn, GormDB: gormDBConn}

	// direct SQL routes
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// gorm routes
	r.HandleFunc("/gorm/users", userHandler.GetGormUsers).Methods("GET")
	r.HandleFunc("/gorm/users", userHandler.CreateGormUser).Methods("POST")
	r.HandleFunc("/gorm/users/{id}", userHandler.UpdateGormUser).Methods("PUT")
	r.HandleFunc("/gorm/users/{id}", userHandler.DeleteGormUser).Methods("DELETE")

	return r
}
