package router

import (
	"github.com/carlosd-ss/go-postgresql/go-postgres/gorila-mux/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", handler.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", handler.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", handler.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id}", handler.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/user/{id}", handler.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}
