package resources

import (
	"log"
	"net/http"

	"github.com/bugjoe/my-cv-backend/persistence"

	"github.com/bugjoe/my-cv-backend/models"
)

// CreateAccountHandler handles account create requests
func CreateAccountHandler(response http.ResponseWriter, request *http.Request) {
	account, err := models.ExtractAccountFromRequest(request)
	if err != nil {
		log.Println("ERROR: Can't extract account from request:", err)
		http.Error(response, "Error while parsing request body", 500)
		return
	}

	log.Println("Create account with email:", account.Email)
	err = persistence.InsertNewAccount(account)
	if err != nil {
		log.Println("ERROR: Can't create new account:", err)
		if err == persistence.ErrAccountAlreadyExists {
			http.Error(response, "User already exists", 403)
		} else {
			http.Error(response, "Error while creating new account", 500)
		}
		return
	}

	response.WriteHeader(200)
}

// GetAccountHandler handles get requests for accounts
func GetAccountHandler(response http.ResponseWriter, request *http.Request) {

}
