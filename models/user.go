package models

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User is mega
type User struct {
	Email    string
	Password string
}

// GetPasswordHash returns the user password as hex encoded SHA-512 hash string
func (user *User) GetPasswordHash() (string, error) {
	sha := crypto.SHA512.New()
	_, err := sha.Write([]byte(user.Password))
	if err != nil {
		return "", err
	}
	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash), nil
}

// UserFromRequest extracts the user from a request
func UserFromRequest(r *http.Request) (*User, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	user := new(User)
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
