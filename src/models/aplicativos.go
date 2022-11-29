package models

import (
	"fmt"
	"proyectoBD/src/database"
)

type Menu struct {
	RoleNegocio string `json:"roleNegocio"`
	Aplicativo  string `json:"aplicativo"`
	Rol         string `json:"rol"`
	Menu        string `json:"menu"`
	RolNegId    int64  `json:"rolNegId"`
	AppRolId    int64  `json:"appRolId"`
	RolId       int64  `json:"rolId"`
	MenuId      int64  `json:"menuId"`
	UserId      int64  `json:"userId"`
	Email       string `json:"email"`
}

func (m *Menu) GetMenu() ([]Menu, error) {
	query := fmt.Sprintf("SELECT * FROM roles_aplicativos_usuario WHERE Email = '%s'", m.Email)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting menu: ", err)
		return nil, err
	}

	var menuList []Menu
	for rows.Next() {
		var m Menu
		err = rows.Scan(&m.RoleNegocio, &m.Aplicativo, &m.Rol, &m.Menu, &m.RolNegId, &m.AppRolId, &m.RolId, &m.MenuId, &m.UserId, &m.Email)
		if err != nil {
			ErrorLogger.Println("Error scanning menu: ", err)
			return nil, err
		}
		menuList = append(menuList, m)
	}

	return menuList, nil
}
