package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"proyectoBD/src/responses"
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

func Routers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	enableCORS(r)

	//api version 1
	v1 := r.PathPrefix("/api/v1").Subrouter()

	auth := v1.PathPrefix("/authenticate").Subrouter()
	signup := v1.PathPrefix("/signup").Subrouter()
	userPreferences := v1.PathPrefix("/userPreferences").Subrouter()
	question := v1.PathPrefix("/questions").Subrouter()

	r.NotFoundHandler = http.HandlerFunc(NotFound)
	InfoLogger.Println("CORS enabled")

	AuthRouter(auth)
	InfoLogger.Println("Auth router enabled at /api/v1/authenticate")

	SignUpRouter(signup)
	InfoLogger.Println("User router enabled at /api/v1/signup")

	UserPreferencesRouter(userPreferences)
	InfoLogger.Println("User preferences router enabled at /api/v1/userPreferences")

	QuestionRouter(question)
	InfoLogger.Println("Question router enabled at /api/v1/questions")

	return r
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses.Exception{Message: "path not found"})
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses.Exception{Message: "method not allowed"})
	}
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
