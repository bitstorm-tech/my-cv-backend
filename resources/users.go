package resources

import (
	"fmt"
	"net/http"

	"github.com/bugjoe/my-cv-backend/models"
)

// UserCreateHandler handles user create requests
func UserCreateHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("UserCreateHandler")
	user, err := models.UserFromRequest(request)
	if err != nil {
		fmt.Println("ERROR: Can't extract user from request")
	}

	fmt.Println("Create user with email:", user.Email)

	response.Write([]byte(fmt.Sprintf("Create user with email: %s", user.Email)))
}

// UserGetHandler handles get requests
func UserGetHandler(response http.ResponseWriter, request *http.Request) {

}
