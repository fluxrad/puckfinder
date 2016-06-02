package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/fluxrad/puckfinder/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Handle API calls
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/rinks", api.RinksHandler)
	s.HandleFunc("/rinks/{id}", api.RinksIDHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
