package resources

import "net/http"

// UserCreateHandler handles user create requests
func UserCreateHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hallo"))
}
