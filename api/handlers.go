package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type Rink struct {
	RinkID    int    `json:"id"`
	ShortName string `json:"shortName"`
	URL       string `json:"url"`
	API       string `json:"api"`
}

type RinkInfo struct {
	RinkID int          `json:"id"`
	Skates []*SkateInfo `json:"skateInfo"`
}

type SkateInfo struct {
	SkateType string `json:"skateType"` // Drop-In, Stick and Puck, etc.
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
}

var Rinks []*Rink

func init() {
	Rinks = []*Rink{
		&Rink{
			RinkID:    1,
			ShortName: "University of Denver",
			URL:       "http://recreation.du.edu",
			API:       "http://denveruniv-web.ungerboeck.com/Calendar/Default.aspx?EventClassFilter=classFilter1&EventFormat=FULLCALENDARJSON",
		},
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
