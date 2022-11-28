package responses

import "proyectoBD/src/models"

// UserPreferencesResponse is a response.
type UserPreferencesResponse struct {
	User models.User `json:"user"`
}
