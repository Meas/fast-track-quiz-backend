package main

import (
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/agg/", indexHandler)
}
