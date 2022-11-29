package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
	"proyectoBD/src/models"
	"strings"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func validateToken(r *http.Request) (models.User, error) {
	//obtener el token desde el header Authorization
	auth := r.Header.Get("Authorization")
	if auth != "" {
		//separar el token del string "Bearer "
		bearerToken := strings.Split(auth, " ")[1]

		// validar el token
		token, _ := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return jwtToken, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var user models.User
			mapstructure.Decode(claims, &user)
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var user models.User
				mapstructure.Decode(claims, &user)

			}
			return user, nil
		} else {
			return models.User{}, fmt.Errorf("Invalid authorization token")
		}
	} else {
		return models.User{}, fmt.Errorf("An authorization header is required")
	}
}
