package persistence

import (
	driver "github.com/arangodb/go-driver"
	"github.com/bugjoe/my-cv-backend/models"
)

// UpsertProfile either inserts a new profile into the database or updates an existing one
func UpsertProfile(profile *models.Profile) (*models.Profile, error) {
	collection, err := getArangoCollection("profiles")
	if err != nil {
		return nil, err
	}

	var meta driver.DocumentMeta

	if len(profile.Key) > 0 {
		_, err = collection.ReplaceDocument(nil, profile.Key, profile.Payload)
	} else {
		meta, err = collection.CreateDocument(nil, profile.Payload)
		profile.Key = meta.Key
	}

	if err != nil {
		return nil, err
	}

	err = createEdge(driver.NewDocumentID("accounts", profile.AccountKey), profile.GetID(), "has")
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// GetProfilesByAccount returns all profiles that are linked to an account.
func GetProfilesByAccount(account models.Account) ([]models.Profile, error) {
	collection, err := getArangoCollection("profiles")
	if err != nil {
		return nil, err
	}

	profiles := []models.Profile{}
	var profile models.Profile

	for _, key := range account.ProfileKeys {
		meta, err := collection.ReadDocument(nil, key, &profile.Payload)
		if err != nil {
			return nil, err
		}
		profile.AccountKey = account.Key
		profile.Key = meta.Key

		profiles = append(profiles, profile)
	}

	return profiles, nil
}
