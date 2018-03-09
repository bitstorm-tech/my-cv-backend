package resources

import (
	"fmt"
	"log"
	"net/http"

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

	response.Write([]byte(fmt.Sprintf("Create account with email: %s", account.Email)))
}

// GetAccountHandler handles get requests for accounts
func GetAccountHandler(response http.ResponseWriter, request *http.Request) {

}
