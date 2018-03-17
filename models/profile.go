package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Profile contains all the main data like title, name, birthday and so on
type Profile struct {
	ID      string `json:"id"`
	Payload struct {
		Title         string `json:"title"`
		AcademicTitle string `json:"academicTitle"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Birthday      string `json:"birthday"`
		Nationality   string `json:"nationality"`
	} `json:"payload"`
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

// ToJSON converts a profile to a JSON []byte
func (profile Profile) ToJSON() ([]byte, error) {
	json, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}

	return json, nil
}
