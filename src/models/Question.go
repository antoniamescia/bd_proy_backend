package models

import (
	"fmt"
	"proyectoBD/src/database"
)

type Question struct {
	IdPregunta int64  `json:"IdPregunta"`
	Pregunta   string `json:"Pregunta"`
}

type PeopleQuestion struct {
	IdPregunta int64  `json:"IdPregunta"`
	UserId     int64  `json:"UserId"`
	Respuesta  string `json:"Respuesta"`
}

func (q *Question) GetQuestions() ([]Question, error) {
	query := fmt.Sprintf("SELECT preg_id as IdPregunta ,pregunta as Pregunta FROM preguntas")
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting questions: ", err)
		return nil, err
	}
	defer rows.Close()
	var questions []Question
	for rows.Next() {
		var question Question
		err := rows.Scan(&question.IdPregunta, &question.Pregunta)
		if err != nil {
			ErrorLogger.Println("Error scanning questions: ", err)
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (q *PeopleQuestion) GetPeopleAnswer() error {
	query := fmt.Sprintf("SELECT preg_id as IdPregunta, user_id as UserId, respuesta as Respuesta FROM personas_preguntas WHERE user_id = %d", q.UserId)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting questions: ", err)
		return err
	}
	var peopleAnswers PeopleQuestion

	for rows.Next() {
		err := rows.Scan(&q.IdPregunta, &q.UserId, &q.Respuesta)
		if err != nil {
			ErrorLogger.Println("Error scanning questions: ", err)
			return err
		}
	}
	fmt.Println("la respuesta desde la base es", peopleAnswers)
	return nil
}

func (q *PeopleQuestion) InsertPeopleAnswer() error {
	query := fmt.Sprintf("INSERT INTO personas_preguntas (preg_id, user_id, respuesta) VALUES (%d, %d, '%s')", q.IdPregunta, q.UserId, q.Respuesta)
	_, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error inserting question answer: ", err)
		return err
	}
	return nil
}
