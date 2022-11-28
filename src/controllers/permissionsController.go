package controllers

import (
	"encoding/json"
	"net/http"
	"proyectoBD/src/models"
	"proyectoBD/src/responses"
)

//type CreateUserResponse struct {
//	ID     int    `json:"id"`
//	Status string `json:"status"`
//}

func GetUserRoles(w http.ResponseWriter, r *http.Request) {
	roles := models.UserRole{}
	email := r.URL.Query().Get("email")
	rolesList, err := roles.GetUserRoles(email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rolesList)
}
