package main

import (
	"github.com/gorilla/mux"
	"github.com/meas/fast-track-quiz-backend/controllers"
)

//InitRoutes - Initialize app routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/questions", controllers.GetQuestions)

	return router
}
