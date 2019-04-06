package main

import (
	"github.com/gorilla/mux"
	"github.com/Meas/fast-track-quiz-backend/controllers"
)

//InitRoutes - Initialize app routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/questions", controllers.GetQuestions).Methods("GET")
	router.HandleFunc("/answers", controllers.ProcessAnswers).Methods("POST")

	return router
}
