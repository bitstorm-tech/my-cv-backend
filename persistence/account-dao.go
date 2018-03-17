package persistence

import (
	"errors"
	"log"

	"github.com/bugjoe/my-cv-backend/models"
)

// ErrAccountAlreadyExists is returned when an account already exists in the database
var ErrAccountAlreadyExists = errors.New("Account already exists")

// // ErrAccountNotFound is returned when no account was found
// var ErrAccountNotFound = errors.New("Account not found")

// InsertNewAccount inserts a new account. When either the Email of the
// account already exists, the functions returns an error
func InsertNewAccount(account *models.Account) error {
	passwordHash, err := account.GetPasswordHash()
	if err != nil {
		return err
	}

	account.Payload.Password = passwordHash
	log.Println("Insert new account:", account)

	database, err := getArangoDatabase()
	if err != nil {
		return err
	}

	query := "FOR acc IN accounts FILTER LOWER(acc.Email) == LOWER(@email) RETURN acc"
	bindings := bindingVariables{
		"email": account.Payload.Email,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return err
	}

	if cursor.HasMore() {
		return ErrAccountAlreadyExists
	}

	collection, err := database.Collection(nil, "accounts")
	if err != nil {
		return err
	}

	_, err = collection.CreateDocument(nil, account.Payload)
	if err != nil {
		return err
	}

	return nil
}

// GetAccountByEmail returns the account that matches with the given email
func GetAccountByEmail(email string) (*models.Account, error) {
	database, err := getArangoDatabase()
	if err != nil {
		return nil, err
	}

	query := "FOR acc IN accounts FILTER LOWER(acc.email) == LOWER(@email) RETURN acc"
	bindings := bindingVariables{
		"email": email,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		// log.Printf("Account with email=%s not found: %v", email, err)
		// return nil, ErrAccountNotFound
		return nil, err
	}

	if !cursor.HasMore() {
		return nil, nil
	}

	account := new(models.Account)
	meta, err := cursor.ReadDocument(nil, &account.Payload)
	if err != nil {
		return nil, err
	}

	account.ID = meta.Key

	if cursor.HasMore() {
		log.Printf("ERROR: found multiple accounts with email %s, will use first one\n", email)
	}

	return account, nil
}
