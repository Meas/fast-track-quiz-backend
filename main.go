package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var someString = "aaaaaaaa"

func main() {
	r := mux.NewRouter()
	InitRoutes(r)
	log.Fatal(http.ListenAndServe(":8081", r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Whoa this is neat!</h1>"+someString)
}
