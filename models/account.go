package models

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Account represents a user account
type Account struct {
	Email    string
	Password string
}

// GetPasswordHash returns the account password as hex encoded SHA-512 hash string
func (account *Account) GetPasswordHash() (string, error) {
	sha := crypto.SHA512.New()
	_, err := sha.Write([]byte(account.Password))
	if err != nil {
		return "", err
	}
	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash), nil
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
