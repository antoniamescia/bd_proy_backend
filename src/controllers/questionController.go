package controllers

import (
	"encoding/json"
	"net/http"
	"proyectoBD/src/models"
	"proyectoBD/src/responses"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := models.Question{}
	questionsList, err := questions.GetQuestions()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questionsList)
}