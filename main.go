package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := InitRoutes()
	log.Fatal(
		http.ListenAndServe(":8081",
			handlers.CORS(handlers.AllowedHeaders([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}))(router),
		))
}
