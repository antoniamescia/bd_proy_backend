package controllers

import (
	"encoding/json"
	"net/http"
	"proyectoBD/src/responses"
)

func GetUserPreferences(w http.ResponseWriter, r *http.Request) {

	user, errToken := validateToken(r)
	response := responses.UserPreferencesResponse{}

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}

	userInfo, errUser := user.GetUser()

	if errUser != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: errUser.Error()})
		return
	}

	response.User = userInfo

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
