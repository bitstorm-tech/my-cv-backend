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
	router.HandleFunc("/accounts", resources.CreateAccountHandler).Methods("PUT")
	router.HandleFunc("/accounts", resources.GetAccountHandler).Methods("GET")
	router.HandleFunc("/profiles", resources.UpsertProfileHandler).Methods("PUT")
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	log.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	log.Println("Server shutdown ...")
}
