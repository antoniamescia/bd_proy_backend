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

func GetPermissionRequests(w http.ResponseWriter, r *http.Request) {

	//validate token
	_, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}

	permissions := models.PermissionRequests{}
	permissionsList, err := permissions.GetPermissionRequests()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(permissionsList)

}

func UpdatePermissionRequest(w http.ResponseWriter, r *http.Request) {

	//validate token
	_, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}

	permissions := models.PermissionRequests{}

	err := json.NewDecoder(r.Body).Decode(&permissions)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	err = permissions.UpdatePermissionRequest()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Response{Data: "Permission request updated successfully"})

}

func CreatePermissionRequest(w http.ResponseWriter, r *http.Request) {

	//validate token
	_, errToken := validateToken(r)

	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(responses.Exception{Message: errToken.Error()})
		return
	}

	permissions := models.PermissionRequests{}

	err := json.NewDecoder(r.Body).Decode(&permissions)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	err = permissions.CreatePermissionRequest()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Response{Data: "Permission request created successfully"})

}

func GetRolesAplicativos(w http.ResponseWriter, r *http.Request) {
	roles := models.RolesAplicativos{}
	rolesList, err := roles.GetRolesAplicativos()
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
