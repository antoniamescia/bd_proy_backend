package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"proyectoBD/src/controllers"
)

func QuestionRouter(r *mux.Router) *mux.Router {
	u := r.PathPrefix("").Subrouter()
	// allow CORS
	u.Use(mux.CORSMethodMiddleware(u))
	u.HandleFunc("", controllers.GetQuestions).Methods("GET")
	u.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return u
}
