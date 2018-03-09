package resources

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bugjoe/my-cv-backend/models"
)

// UserCreateHandler handles user create requests
func UserCreateHandler(response http.ResponseWriter, request *http.Request) {
	user, err := models.UserFromRequest(request)
	if err != nil {
		log.Println("ERROR: Can't extract user from request:", err)
		http.Error(response, "Error while parsing request body", 500)
		return
	}

	log.Println("Create user with email:", user.Email)

	response.Write([]byte(fmt.Sprintf("Create user with email: %s", user.Email)))
}

// UserGetHandler handles get requests
func UserGetHandler(response http.ResponseWriter, request *http.Request) {

}
