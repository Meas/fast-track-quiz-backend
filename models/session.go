package models

// Session - structure for quiz session
type Session struct {
	ID        string     `json:"id"`
	Questions []Question `json:"questions"`
}

// SessionSubmission - structure for quiz submission
type SessionSubmission struct {
	ID      string            `json:"id"`
	Answers []SubmittedAnswer `json:"answers"`
}
