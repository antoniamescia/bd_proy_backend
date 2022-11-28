package controllers

import (
	"encoding/json"
	"net/http"
	"proyectoBD/src/models"
	"proyectoBD/src/responses"
	"strings"
)

type CreateUserResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	type input struct {
		User     models.User           `json:"user"`
		Question models.PeopleQuestion `json:"question"`
	}
	user := models.User{}
	peopleAnswer := models.PeopleQuestion{}
	inputDecode := input{}
	err := json.NewDecoder(r.Body).Decode(&inputDecode)
	//errQuestion := json.NewDecoder(r.Body).Decode(&peopleAnswer)
	user = inputDecode.User
	peopleAnswer = inputDecode.Question

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	id, errCreate := user.CreateUser()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: errCreate.Error()})
		return
	}

	if id == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: errCreate.Error()})
		return
	}

	peopleAnswer.UserId = id
	errQuestionAnswer := peopleAnswer.InsertPeopleAnswer()
	if errQuestionAnswer != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: errQuestionAnswer.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{ID: int(id), Status: "User created"})
	return
}

func RecoverPassword(w http.ResponseWriter, r *http.Request) {

	type input struct {
		Email      string `json:"email"`
		PreguntaID int64  `json:"pregunta_id"`
		Respuesta  string `json:"respuesta"`
	}
	inputDecode := input{}
	err := json.NewDecoder(r.Body).Decode(&inputDecode)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}

	user := models.User{}
	user.Email = inputDecode.Email
	user.GetUser()

	//check if input answer is correct
	peopleAnswer := models.PeopleQuestion{}

	peopleAnswer.UserId = user.UserId

	peopleAnswer.GetPeopleAnswer()

	if strings.ToLower(peopleAnswer.Respuesta) != strings.ToLower(inputDecode.Respuesta) || peopleAnswer.IdPregunta != inputDecode.PreguntaID {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Incorrect answer"})
		return
	}

	//update password
	pwd, errpwd := user.GenerateNewPassword()
	if errpwd != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.Exception{Message: errpwd.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Response{pwd})
	return
}
