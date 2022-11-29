package controllers

import (
	"encoding/json"
	"net/http"
	"proyectoBD/src/models"
	"proyectoBD/src/responses"
)

//get menu for user

func GetMenu(w http.ResponseWriter, r *http.Request) {
	//validate token
	user, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}

	//get menu
	menu := models.Menu{}
	menu.Email = user.Email
	menuList, err := menu.GetMenu()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menuList)
}
