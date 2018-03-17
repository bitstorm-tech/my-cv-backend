package resources

import (
	"log"
	"net/http"

	"github.com/bugjoe/my-cv-backend/persistence"

	"github.com/bugjoe/my-cv-backend/models"
)

// UpsertProfileHandler handles insert or update (aka upsert) requests for profiles
func UpsertProfileHandler(response http.ResponseWriter, request *http.Request) {
	profile, err := models.ExtractProfileFromRequest(request)
	if err != nil {
		http.Error(response, "Error while parsing body", 500)
		return
	}

	profile, err = persistence.UpsertProfile(profile)
	if err != nil {
		log.Println("ERROR:", err)
		http.Error(response, "Error wile upserting profile", 500)
		return
	}

	json, err := profile.ToJSON()
	if err != nil {
		log.Println("ERROR:", err)
		http.Error(response, "Error while create JSON", 500)
		return
	}
	response.Write(json)
}
