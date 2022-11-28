package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"proyectoBD/src/controllers"
)

func UserPreferencesRouter(r *mux.Router) *mux.Router {
	u := r.PathPrefix("").Subrouter()
	// allow CORS
	u.Use(mux.CORSMethodMiddleware(u))
	u.HandleFunc("", controllers.GetUserPreferences).Methods("GET")
	u.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return u
}
