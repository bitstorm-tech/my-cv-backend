package persistence

import (
	"log"

	"github.com/bugjoe/my-cv-backend/models"
)

// InsertNewAccount inserts a new account. When either the Email of the
// account already exists, the functions returns an error
func InsertNewAccount(account *models.Account) error {
	passwordHash, err := account.GetPasswordHash()
	if err != nil {
		return err
	}

	account.Password = passwordHash
	log.Println("Insert new account:", account)

	database := *getArangoDatabase()
	query := "FOR acc IN accounts FILTER LOWER(acc.Email) == LOWER(@email) RETURN acc"
	bindings := bindingVariables{
		"email": account.Email,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return err
	}

	if cursor.HasMore() {
		return err
	}

	collection, err := database.Collection(nil, "accounts")
	if err != nil {
		return err
	}

	_, err = collection.CreateDocument(nil, account)
	if err != nil {
		return err
	}

	return nil
}

// GetAccountByEmail returns the account that matches with the given email
func GetAccountByEmail(email string) (*models.Account, error) {
	log.Println("Get account by email:", email)

	database := *getArangoDatabase()
	query := "FOR acc IN accounts FILTER LOWER(acc.Email) == LOWER(@email) RETURN u"
	bindings := bindingVariables{
		"email": email,
	}
	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return nil, err
	}

	account := new(models.Account)
	_, err = cursor.ReadDocument(nil, account)
	if err != nil {
		return nil, err
	}

	if cursor.HasMore() {
		log.Printf("ERROR: found multiple accounts with email %s, will use first one\n", email)
	}

	return account, nil
}
