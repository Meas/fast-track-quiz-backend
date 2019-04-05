package models

type Session struct {
	ID        string     `json:"id"`
	Questions []Question `json:"questions"`
}

type SessionSubmission struct {
	ID      string            `json:"id"`
	Answers []SubmittedAnswer `json:"answers"`
}
