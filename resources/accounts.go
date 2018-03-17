package resources

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/bugjoe/my-cv-backend/persistence"

	"github.com/bugjoe/my-cv-backend/models"
)

var secret = []byte("OSR6eqMv7N01PlPFZyBS1k508daeP8hC15dwRQ5pzr7hwsIOcAWQuhdZlGUKHIQw")

// CreateAccountHandler handles account create requests
func CreateAccountHandler(response http.ResponseWriter, request *http.Request) {
	account, err := models.ExtractAccountFromRequest(request)
	if err != nil {
		log.Println("ERROR: Can't extract account from request:", err)
		http.Error(response, "Error while parsing request body", 500)
		return
	}

	log.Println("Create account with email:", account.Payload.Email)
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

// LoginHandler handles login requests. If the login is successfull, it will
// respond with a JWT.
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	account, err := models.ExtractAccountFromRequest(request)
	if err != nil {
		log.Println("ERROR: Can't extract account from request", err)
		http.Error(response, "Error while parsing request body", 500)
		return
	}

	accountFromDb, err := persistence.GetAccountByEmail(account.Payload.Email)
	if err != nil {
		log.Println("ERROR: Can't get account from database", err)
		http.Error(response, "Error while getting account from database", 500)
		return
	}

	if accountFromDb == nil {
		http.Error(response, "Invalid login", 401)
		return
	}

	passwordHash, err := account.GetPasswordHash()
	if err != nil {
		log.Println("ERROR: Can't get password hash from account", err)
		http.Error(response, "Error while creating password hash", 500)
		return
	}

	if passwordHash != accountFromDb.Payload.Password {
		http.Error(response, "Invalid login", 401)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		ExpiresAt: 60000,
		Id:        accountFromDb.Payload.Email,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println("ERROR: Can't sign JWT", err)
		http.Error(response, "Error wile signing JWT", 500)
		return
	}

	response.Write([]byte(tokenString))
}
