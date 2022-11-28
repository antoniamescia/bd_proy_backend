package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"proyectoBD/src/controllers"
)

func AuthRouter(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.CreateToken).Methods("POST")
	a.HandleFunc("/recoverpwd", controllers.RecoverPassword).Methods("POST")

	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}
