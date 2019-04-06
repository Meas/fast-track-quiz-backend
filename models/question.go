package models

import (
	"github.com/Meas/fast-track-quiz-backend/helpers"
)

// Question - structure for a single question
type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

// GetQuestions - loads questions from json file
func GetQuestions() []Question {
	var questions = make([]Question, 0)
	helpers.LoadJSON("questions.json", &questions)
	return questions
}
