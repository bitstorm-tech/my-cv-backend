package persistence

import (
	"log"

	"github.com/bugjoe/my-cv-backend/models"
)

// InsertNewUser inserts a new user. When either the Username or the Email of the
// user already exists, the functions returns an error
func InsertNewUser(user *models.User) error {
	passwordHash, err := user.GetPasswordHash()
	if err != nil {
		return err
	}

	user.Password = passwordHash
	log.Println("Insert new user:", user)

	database := *getArangoDatabase()
	query := "FOR u IN users FILTER LOWER(u.Email) == LOWER(@email) RETURN u"
	bindings := bindingVariables{
		"email": user.Email,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return err
	}

	if cursor.HasMore() {
		return err
	}

	collection, err := database.Collection(nil, "users")
	if err != nil {
		return err
	}

	_, err = collection.CreateDocument(nil, user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail returns the user that matches with the given email
func GetUserByEmail(email string) (*models.User, error) {
	log.Println("Get user by email:", email)

	database := *getArangoDatabase()
	query := "FOR u IN users FILTER LOWER(u.Email) == LOWER(@email) RETURN u"
	bindings := bindingVariables{
		"email": email,
	}
	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return nil, err
	}

	user := new(models.User)
	_, err = cursor.ReadDocument(nil, user)
	if err != nil {
		return nil, err
	}

	if cursor.HasMore() {
		log.Printf("ERROR: found multiple users with email %s, will use first one\n", email)
	}

	return user, nil
}
