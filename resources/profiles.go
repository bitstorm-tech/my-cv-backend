package resources

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bugjoe/my-cv-backend/models"
)

// CreateProfileHandler handles create (PUT) requests for profiles
func CreateProfileHandler(response http.ResponseWriter, request *http.Request) {
	profile, err := models.ExtractProfileFromRequest(request)
	if err != nil {
		http.Error(response, "Error while parsing body", 500)
		return
	}

	fmt.Println("GOT PROFILE!!!!!!!!!", profile)
	profile.ID = fmt.Sprintf("%d", time.Now().Nanosecond())
	json, err := profile.ToJSON()
	if err != nil {
		http.Error(response, "Error while create JSON", 500)
		return
	}
	response.Write(json)
}
