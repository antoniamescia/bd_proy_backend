package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"proyectoBD/src/controllers"
)

func PermissionRouter(r *mux.Router) *mux.Router {
	u := r.PathPrefix("").Subrouter()
	// allow CORS
	u.Use(mux.CORSMethodMiddleware(u))
	u.HandleFunc("", controllers.GetUserRoles).Methods("GET")
	u.HandleFunc("/requests", controllers.GetPermissionRequests).Methods("GET")
	u.HandleFunc("/requests", controllers.UpdatePermissionRequest).Methods("PUT")
	u.HandleFunc("/requests", controllers.CreatePermissionRequest).Methods("POST")
	u.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return u
}
