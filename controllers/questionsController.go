package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meas/fast-track-quiz-backend/helpers"
	"github.com/meas/fast-track-quiz-backend/models"
)

type Session struct {
	ID        string            `json:"id"`
	Questions []models.Question `json:"question"`
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	response := Session{
		ID:        helpers.GetRandomString(10),
		Questions: models.GetQuestions(),
	}
	fmt.Println(response)
	json.NewEncoder(w).Encode(response)
}
