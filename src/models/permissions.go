package models

import (
	"fmt"
	"proyectoBD/src/database"
)

type permissionRequests struct {
	Nombre            string `json:"nombre"`
	Email             string `json:"email"`
	Aplicacion        string `json:"aplicacion"`
	Permiso           string `json:"permiso"`
	FechaSolicitud    string `json:"fechaSolicitud"`
	FechaAutorizacion string `json:"fechaAutorizacion"`
	Estado            string `json:"estado"`
	UserId            int64  `json:"userId"`
	AppId             int64  `json:"appId"`
	RolNegId          int64  `json:"rolNegId"`
}

type UserRole struct {
	RolId       int64  `json:"rolId"`
	Description string `json:"description"`
}

//GetPermissionRequests returns all the permission requests.
func (p *permissionRequests) GetPermissionRequests() ([]permissionRequests, error) {
	query := fmt.Sprintf("SELECT * FROM permisos_solicitados")
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting permission requests: ", err)
		return nil, err
	}

	var permissionRequestsList []permissionRequests
	for rows.Next() {
		var p permissionRequests
		err = rows.Scan(&p.Nombre, &p.Email, &p.Aplicacion, &p.Permiso, &p.FechaSolicitud, &p.FechaAutorizacion, &p.Estado, &p.UserId, &p.AppId, &p.RolNegId)
		if err != nil {
			ErrorLogger.Println("Error scanning permission requests: ", err)
			return nil, err
		}
		permissionRequestsList = append(permissionRequestsList, p)
	}

	return permissionRequestsList, nil
}

//Get roles of a user
func (u *UserRole) GetUserRoles(UserEmail string) ([]UserRole, error) {
	query := fmt.Sprintf("SELECT rol_neg_id as RolId, descripcion_rol_neg as Descripcion FROM roles_usuario WHERE email = '%s'", UserEmail)
	fmt.Println(query)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting user roles: ", err)
		return nil, err
	}

	var userRolesList []UserRole
	for rows.Next() {
		var u UserRole
		err = rows.Scan(&u.RolId, &u.Description)
		if err != nil {
			ErrorLogger.Println("Error scanning user roles: ", err)
			return nil, err
		}
		userRolesList = append(userRolesList, u)
	}
	fmt.Println(userRolesList)
	return userRolesList, nil
}
