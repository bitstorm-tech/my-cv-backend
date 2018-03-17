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

	if len(profile.ID) > 0 {
		_, err = collection.ReplaceDocument(nil, profile.ID, profile.Payload)
	} else {
		meta, err = collection.CreateDocument(nil, profile.Payload)
		profile.ID = meta.Key
	}

	if err != nil {
		return nil, err
	}

	return profile, nil
}
