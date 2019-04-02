package main

import (
	"log"
	"net/http"
)

func main() {
	router := InitRoutes()
	log.Fatal(http.ListenAndServe(":8081", router))
}
