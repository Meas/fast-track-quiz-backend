package models

import (
	"github.com/meas/fast-track-quiz-backend/helpers"
)

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

//var questions map[int]Question

func GetQuestions() []Question {
	var questions = make([]Question, 0)
	helpers.LoadJSON("questions.json", &questions)
	return questions
}
