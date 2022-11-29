package models

import (
	"fmt"
	"proyectoBD/src/database"
)

type PermissionRequests struct {
	Nombres           string `json:"nombre"`
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

type Aplicativos struct {
	AppId     int64  `json:"appId"`
	NombreApp string `json:"nombreApp"`
}

type RolesAplicativos struct {
	RolNegId    int64         `json:"rolNegId"`
	Descripcion string        `json:"descripcion"`
	Aplicativos []Aplicativos `json:"aplicativos"`
}

func (p *RolesAplicativos) GetRolesAplicativos() ([]RolesAplicativos, error) {
	query := fmt.Sprintf("select rol_neg_id as RolNegID, descripcion_rol_neg as Descripcion from roles_negocio")
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting roles aplicativos: ", err)
		return nil, err
	}

	var rolesAplicativosList []RolesAplicativos
	for rows.Next() {
		var p RolesAplicativos
		err = rows.Scan(&p.RolNegId, &p.Descripcion)
		if err != nil {
			ErrorLogger.Println("Error scanning roles aplicativos: ", err)
			return nil, err
		}
		rolesAplicativosList = append(rolesAplicativosList, p)
	}
	fmt.Println(rolesAplicativosList)

	for i, _ := range rolesAplicativosList {
		query := fmt.Sprintf("select AppId, NombreApp from nombre_aplicativos_roles where RolNegID = %d", rolesAplicativosList[i].RolNegId)
		rows, err := database.QueryDB(query)
		if err != nil {
			ErrorLogger.Println("Error getting aplicativos: ", err)
			return nil, err
		}
		var aplicativosList []Aplicativos
		for rows.Next() {
			var p Aplicativos
			err = rows.Scan(&p.AppId, &p.NombreApp)
			if err != nil {
				ErrorLogger.Println("Error scanning aplicativos: ", err)
				return nil, err
			}
			aplicativosList = append(aplicativosList, p)
		}
		rolesAplicativosList[i].Aplicativos = aplicativosList
	}

	return rolesAplicativosList, nil
}

//GetPermissionRequests returns all the permission requests.
func (p *PermissionRequests) GetPermissionRequests() ([]PermissionRequests, error) {
	query := fmt.Sprintf("SELECT * FROM permisos_solicitados")
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting permission requests: ", err)
		return nil, err
	}

	var permissionRequestsList []PermissionRequests
	for rows.Next() {
		var p PermissionRequests
		err = rows.Scan(&p.Nombres, &p.Email, &p.Aplicacion, &p.Permiso, &p.FechaSolicitud, &p.FechaAutorizacion, &p.Estado, &p.UserId, &p.AppId, &p.RolNegId)
		if err != nil {
			ErrorLogger.Println("Error scanning permission requests: ", err)
			return nil, err
		}
		permissionRequestsList = append(permissionRequestsList, p)
	}

	return permissionRequestsList, nil
}

//update permission request
func (p *PermissionRequests) UpdatePermissionRequest() error {
	query := fmt.Sprintf("UPDATE permisos SET estado = '%s', fecha_autorizacion = now() WHERE user_id = %d AND app_id = %d AND rol_neg_id = %d", p.Estado, p.UserId, p.AppId, p.RolNegId)
	_, err := database.UpdateDB(query)
	if err != nil {
		ErrorLogger.Println("Error updating permission request: ", err)
		return err
	}
	return nil
}

//Get roles of a user
func (u *UserRole) GetUserRoles(UserEmail string) ([]UserRole, error) {
	query := fmt.Sprintf("SELECT rol_neg_id as RolId, descripcion_rol_neg as Descripcion FROM roles_usuario WHERE email = '%s' order by rol_neg_id asc", UserEmail)
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

//Create a new permission request
func (p *PermissionRequests) CreatePermissionRequest() error {
	query := fmt.Sprintf("INSERT INTO permisos(user_id, app_id, rol_neg_id, estado, fecha_solicitud,fecha_autorizacion) VALUES (%d, %d, %d,'PENDIENTE', now(), now())", p.UserId, p.AppId, p.RolNegId)

	_, err := database.UpdateDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting permission requests: ", err)
		return err
	}
	return nil
}
