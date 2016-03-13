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
	s.HandleFunc("/rinks/{rinkID}/", api.RinkIDHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
