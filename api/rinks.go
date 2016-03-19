package api

import (
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/patrickmn/go-cache"
)

// Store this in a database eventually
type Rink struct {
	RinkID    int    `json:"id"`
	ShortName string `json:"shortName"`
	URL       string `json:"url"`
	API       string `json:"api"`
	Parser    string `json:"parser"`
	Timeout   int    `json:"timeout"`
}

type Skate struct {
	Type      string `json:"type"` // Drop-In, Stick and Puck, etc.
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
}

// We need a getter to setup the thing.
func (r *Rink) Skates() []*Skate {
	skates, found := Cache.Get(strconv.Itoa(r.RinkID))
	if found {
		return skates.([]*Skate)
	}

	skates, err := fetchSkateData(r.API, r.Parser)
	if err != nil {
		log.Error(err)
	}
	Cache.Set(strconv.Itoa(r.RinkID), skates, cache.DefaultExpiration)

	return skates.([]*Skate)
}

func fetchSkateData(api string, parser string) ([]*Skate, error) {
	var skates []*Skate
	var p SkateParser

	resp, err := http.Get(api)
	if err != nil {
		log.Error("Could not retrieve data: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Could not read body for %s: %s", api, err)
	}

	switch parser {
	case `calendarjson`:
		p = &CalendarJSON{}
	}

	skates, err = p.Parse(body)
	if err != nil {
		log.Error("Could not parse skates: %s", err)
	}

	return skates, nil
}
