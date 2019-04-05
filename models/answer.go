package models

import (
	"github.com/meas/fast-track-quiz-backend/helpers"
)

type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type SubmittedAnswer struct {
	QuestionID int `json:"questionId"`
	AnswerID   int `json:"answerId"`
}

func GetCorrectAnswers() []SubmittedAnswer {
	var answers = make([]SubmittedAnswer, 0)
	helpers.LoadJSON("answers.json", &answers)
	return answers
}
