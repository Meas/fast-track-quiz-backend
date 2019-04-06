package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Meas/fast-track-quiz-backend/helpers"
	"github.com/Meas/fast-track-quiz-backend/models"
	"github.com/Meas/fast-track-quiz-backend/repositories"
)

// So everything is in memory
var allQuestions = models.GetQuestions()
var correctAnswers = models.GetCorrectAnswers()
var allSubmissions = make([]models.Score, 0)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	response := models.Session{
		ID:        helpers.GetRandomString(10),
		Questions: allQuestions,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ProcessAnswers(w http.ResponseWriter, r *http.Request) {
	/* decoder := json.NewDecoder(r.Body)
	json.NewEncoder(w).Encode(decoder) */
	var submittedSession models.SessionSubmission

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &submittedSession)

	var result = repositories.GetQuizResult(submittedSession, correctAnswers, &allSubmissions)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
