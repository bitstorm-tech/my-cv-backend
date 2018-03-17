package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	driver "github.com/arangodb/go-driver"
)

// Profile contains all the main data like title, name, birthday and so on
type Profile struct {
	Key     string `json:"key"`
	Payload struct {
		Title         string `json:"title"`
		AcademicTitle string `json:"academicTitle"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Birthday      string `json:"birthday"`
		Nationality   string `json:"nationality"`
	} `json:"payload"`
	AccountKey string `json:"accountKey"`
}

// NewProfile creates a new empty profile
func NewProfile() *Profile {
	profile := new(Profile)
	profile.Key = ""
	profile.Payload.Title = ""
	profile.Payload.AcademicTitle = ""
	profile.Payload.FirstName = ""
	profile.Payload.LastName = ""
	profile.Payload.Birthday = ""
	profile.Payload.Nationality = ""
	profile.AccountKey = ""

	return profile
}

// ExtractProfileFromRequest extracts the profile from a request
func ExtractProfileFromRequest(r *http.Request) (*Profile, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	profile := new(Profile)
	err = json.Unmarshal(body, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// ToJSON converts a Profile to a JSON []byte
func (profile Profile) ToJSON() ([]byte, error) {
	json, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// GetID returns the ID of the profile in the form of "profiles/profile-key"
func (profile *Profile) GetID() driver.DocumentID {
	return driver.NewDocumentID("profiles", profile.Key)
}
