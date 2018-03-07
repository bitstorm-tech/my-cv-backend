package main

import (
	"fmt"
	"my-cv-backend/resources"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start My-CV server")
	router := mux.NewRouter()
	router.HandleFunc("/users", resources.UserCreateHandler).Methods("GET")
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	fmt.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	fmt.Println("Server shutdown ...")
}
