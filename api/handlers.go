package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func AllRinksHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(Rinks)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(resp)
}

func RinkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	i, err := strconv.Atoi(vars["rinkID"])
	if err != nil {
		log.Error("Couldn't find string: %s", err)
		// Write an error here
	}

	// Eventually fetch this from a DB or something
	rink := NewRink(i)

	resp, err := json.Marshal(rink)
	if err != nil {
		log.Error("Couldn't marshal rinks response: %s", err)
		// Write an error here too
	}

	w.Write(resp)
}
