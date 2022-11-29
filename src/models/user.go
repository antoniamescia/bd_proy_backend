package models

import (
	"crypto/rand"
	"fmt"
	"proyectoBD/src/database"
	"proyectoBD/src/hashing"
)

type IUser interface {
	GetUser(username string) User
}

// User is a user.
type User struct {
	UserId       int64  `json:"UserId"`
	Nombres      string `json:"Nombres"`
	Apellidos    string `json:"Apellidos"`
	Direccion    string `json:"Direccion"`
	Ciudad       string `json:"Ciudad"`
	Departamento string `json:"Departamento"`
	Email        string `json:"Email"`
	Password     string `json:"Password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//
func (u *Login) checkUserExists() bool {
	query := fmt.Sprintf("SELECT email as Email FROM personas WHERE email = '%s'", u.Email)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error checking if user exists: ", err)
	}
	i := 0
	var EmailBD string
	for rows.Next() {
		i++
		err = rows.Scan(&EmailBD)
	}
	switch i {
	case 0:
		return false
	default:
		if u.Email == EmailBD {
			return true
		} else {
			return false
		}
	}
}

func (u *Login) ValidateLogin() (bool, error) {
	if !u.checkUserExists() {
		return false, fmt.Errorf("User does not exist")
	}

	query := fmt.Sprintf("SELECT email as Email,hashpwd as Password FROM personas WHERE email = '%s'", u.Email)
	rows, err := database.QueryDB(query)
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	var hashFromBD string
	var emailFromBD string
	for rows.Next() {
		i++
		err = rows.Scan(&emailFromBD, &hashFromBD)
		if err != nil {
			fmt.Println(err)
		}
	}
	// si obtengo 0 resultados, no existe el usuario
	if i == 0 {
		WarningLogger.Println("Error validate login: ", u.Email)
		return false, fmt.Errorf("Password error")
	}
	// si obtengo 1 resultado, valido el hash
	if i == 1 {
		if u.validateHash(hashFromBD) {
			return true, nil
		} else {
			WarningLogger.Println("Error validate login: ", u.Email)
			return false, fmt.Errorf("Password error")
		}
	}
	ErrorLogger.Println("Multiple users with Email: ", u.Email)
	return false, fmt.Errorf("Multiple users with Email: %s", u.Email)

}

func (u *User) CreateUser() (int64, error) {
	uLogin := Login{Email: u.Email}
	if uLogin.checkUserExists() {
		return 0, fmt.Errorf("User already exists")
	}
	pswHashed, errHash := hashing.HashPassword(u.Password)
	if errHash != nil {
		return 0, fmt.Errorf("Error hashing password")
	}
	u.Password = pswHashed
	query := fmt.Sprintf("INSERT INTO personas (nombres, apellidos, direccion, ciudad, departamento, hashpwd, email) VALUES ('%s', '%s', '%s', '%s', '%s', '%s','%s')", u.Nombres, u.Apellidos, u.Direccion, u.Ciudad, u.Departamento, pswHashed, u.Email)
	id, err := database.InsertDB(query)
	if err != nil {
		ErrorLogger.Println("Error creating user: ", err)
		return 0, fmt.Errorf("Error creating user: ", err)
	}
	return id, nil
}

func (u *User) GetUser() (User, error) {

	query := fmt.Sprintf("SELECT user_id as UserId, nombres as Nombres, apellidos as Apellidos, direccion as Direccion, ciudad as Ciudad, departamento as Departamento, email as Email FROM personas WHERE email = '%s'", u.Email)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting user: ", err)
		return User{}, fmt.Errorf("Error getting user: ", err)
	}
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.Nombres, &u.Apellidos, &u.Direccion, &u.Ciudad, &u.Departamento, &u.Email)
		if err != nil {
			ErrorLogger.Println("Error getting user: ", err)
			return User{}, fmt.Errorf("Error getting user: ", err)
		}
	}

	return *u, nil
}

func (u *User) GenerateNewPassword() (string, error) {
	psw := randomString()
	pswHashed, errHash := hashing.HashPassword(psw)
	if errHash != nil {
		return "", fmt.Errorf("Error hashing password")
	}
	u.Password = pswHashed
	query := fmt.Sprintf("UPDATE personas SET hashpwd = '%s' WHERE email = '%s'", u.Password, u.Email)
	_, err := database.UpdateDB(query)
	if err != nil {
		ErrorLogger.Println("Error updating password: ", err)
		return "", fmt.Errorf("Error updating password: ", err)
	}
	return psw, nil
}

func (u *Login) validateHash(hashFromBD string) bool {
	return hashing.CheckPasswordHash(u.Password, hashFromBD)
}

func randomString() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
