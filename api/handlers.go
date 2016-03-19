package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// RinksReponse is a response object containing an array of Rink metadata.
type RinksResponse struct {
	Rinks []*Rink `json:"rinks"`
}

// RinkIDResponse contains information about a specific rink and its skates
type RinksIDResponse struct {
	*Rink  `json:"rink"`
	Skates []*Skate `json:"skates"`
}

// RinksHandler returns a RinksResponse object, containing metadata bout all
// Rinks in the database.
func RinksHandler(w http.ResponseWriter, r *http.Request) {
	// convert to a slice rather than a map so we can marshal it
	var rinks []*Rink
	for _, v := range Rinks {
		rinks = append(rinks, v)
	}

	rr := &RinksResponse{
		Rinks: rinks,
	}

	resp, err := json.Marshal(rr)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(resp)
}

// RinksIDHandler writes a JSON response object containing Rink and Skate
// information specific to a particular RinkID
func RinksIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("Couldn't find string: %s", err)
		// Write an error here
	}

	// Eventually fetch this from a DB
	rink := Rinks[i]

	rir := &RinksIDResponse{
		Rink:   rink,
		Skates: rink.Skates(),
	}

	resp, err := json.Marshal(rir)
	if err != nil {
		log.Error("Couldn't marshal rinks response: %s", err)
		// Write an error here too
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(resp)
}
