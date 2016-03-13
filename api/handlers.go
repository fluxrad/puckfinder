package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type Rink struct {
	RinkID int    `json:"id"`
	URL    string `json:"url"`
	APIURL string `json:"apiUrl"`
}

var Rinks []*Rink

func init() {
	Rinks = []*Rink{
		&Rink{1, "http://recreation.du.edu", "http://denveruniv-web.ungerboeck.com/Calendar/Default.aspx?EventClassFilter=classFilter1&EventFormat=FULLCALENDARJSON"},
	}
}

func RinksHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(Rinks)
	if err != nil {
		log.Error(err)
	}

	w.Write(resp)
}

func RinkIDHandler(w http.ResponseWriter, r *http.Request) {
}
