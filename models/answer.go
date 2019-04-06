package models

import (
	"github.com/Meas/fast-track-quiz-backend/helpers"
)

// Answer - structure for a single answer
type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// SubmittedAnswer - structure for a single submitted answer
type SubmittedAnswer struct {
	QuestionID int `json:"questionId"`
	AnswerID   int `json:"answerId"`
}

// GetCorrectAnswers - loads correct answers from json file
func GetCorrectAnswers() []SubmittedAnswer {
	var answers = make([]SubmittedAnswer, 0)
	helpers.LoadJSON("answers.json", &answers)
	return answers
}
