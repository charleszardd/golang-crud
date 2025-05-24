package main

import (
	"go-api/config"
	"go-api/routes"
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

)

func main(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	config.ConnectDB()

	router := mux.NewRouter()

	routes.RegisterUserRoutes(router)
	// fmt.Println("🔗 Registered Routes:")
	// router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	path, _ := route.GetPathTemplate()
	// 	methods, _ := route.GetMethods()
	// 	fmt.Printf("-> %s %s \n", methods, path)
	// 	return nil
	// })

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

