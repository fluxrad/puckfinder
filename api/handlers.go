package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type RinksResponse struct {
	Rinks []*Rink `json:"rinks"`
}

type RinksIDResponse struct {
	Rink   `json:"rink"`
	Skates []*Skate `json:"skates"`
}

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
	w.Write(resp)
}

func RinksIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error("Couldn't find string: %s", err)
		// Write an error here
	}

	// Eventually fetch this from a DB or something
	rink := *Rinks[i]
	rir := &RinksIDResponse{
		Rink:   rink,
		Skates: rink.Skates(),
	}

	resp, err := json.Marshal(rir)
	if err != nil {
		log.Error("Couldn't marshal rinks response: %s", err)
		// Write an error here too
	}

	w.Write(resp)
}
