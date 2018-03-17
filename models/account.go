package models

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	driver "github.com/arangodb/go-driver"
)

// Account represents a user account
type Account struct {
	Key     string `json:"key"`
	Payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"payload"`
	ProfileKeys []string `json:"profileKeys"`
}

// NewAccount creates a new empty account
func NewAccount() *Account {
	account := new(Account)
	account.Key = ""
	account.Payload.Email = ""
	account.Payload.Password = ""
	account.ProfileKeys = make([]string, 0)

	return account
}

// GetPasswordHash returns the account password as hex encoded SHA-512 hash string
func (account *Account) GetPasswordHash() (string, error) {
	sha := crypto.SHA512.New()
	_, err := sha.Write([]byte(account.Payload.Password))
	if err != nil {
		return "", err
	}
	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash), nil
}

// ToJSON converts an Account to a JSON []byte
func (account *Account) ToJSON() ([]byte, error) {
	json, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// ExtractAccountFromRequest extracts the account from a request
func ExtractAccountFromRequest(r *http.Request) (*Account, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	account := new(Account)
	err = json.Unmarshal(body, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// GetID returns the ID of the account in the form of "accounts/account-key"
func (account *Account) GetID() driver.DocumentID {
	return driver.NewDocumentID("accounts", account.Key)
}
