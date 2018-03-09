package main

import (
	"log"
	"net/http"

	"github.com/bugjoe/my-cv-backend/resources"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Start My-CV server")
	router := mux.NewRouter()
	router.HandleFunc("/users", resources.UserCreateHandler).Methods("PUT")
	router.HandleFunc("/users", resources.UserGetHandler).Methods("GET")
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	log.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	log.Println("Server shutdown ...")
}
