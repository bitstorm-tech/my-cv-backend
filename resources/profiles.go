package resources

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bugjoe/my-cv-backend/persistence"

	"github.com/bugjoe/my-cv-backend/models"
)

// UpsertProfileHandler handles insert or update (aka upsert) requests for profiles
func UpsertProfileHandler(response http.ResponseWriter, request *http.Request) {
	profile, err := models.ExtractProfileFromRequest(request)
	if err != nil {
		log.Println("ERROR:", err)
		http.Error(response, "Error while parsing body", 500)
		return
	}

	profile, err = persistence.UpsertProfile(profile)
	if err != nil {
		log.Printf("ERROR: Can't upsert profile: %v\n%v\n", profile, err)
		http.Error(response, "Error wile upserting profile", 500)
		return
	}

	json, err := profile.ToJSON()
	if err != nil {
		log.Printf("ERROR: Can't convert profile=%v to JSON\n%v\n", profile, err)
		http.Error(response, "Error while create JSON", 500)
		return
	}

	response.Write(json)
}

// GetAllProfilesHandler returns all profiles that are linked to a account. Therefore,
// the email of the account has to be included in the URL (/accounts/email).
func GetAllProfilesHandler(response http.ResponseWriter, request *http.Request) {
	email := mux.Vars(request)["email"]
	account, err := persistence.GetAccountByEmail(email)
	if err != nil {
		log.Printf("ERROR: Can't get account with email=%s from databae\n%v\n", email, err)
		http.Error(response, "Error while getting account from database", 500)
		return
	}

	profiles, err := persistence.GetProfilesByAccount(*account)
	if err != nil {
		log.Printf("ERROR: Can't get profile from account=%v\n%v\n", account, err)
		http.Error(response, "Error while getting profiles from database", 500)
		return
	}

	json, err := json.Marshal(profiles)
	if err != nil {
		log.Printf("ERROR: Can't convert profiles=%v to JSON\n%v\n", profiles, err)
		http.Error(response, "Error while create JSON", 500)
		return
	}

	response.Write(json)
}
