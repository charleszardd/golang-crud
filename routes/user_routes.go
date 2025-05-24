package routes

import (
	"github.com/gorilla/mux"
	"go-api/handlers"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/new-user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", handlers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")
}