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
	RinkID    int     `json:"id"`
	ShortName string  `json:"shortName"`
	URL       string  `json:"url"`
	API       string  `json:"api"`
	Parser    string  `json:"-"`
	Timeout   int     `json:"timeout"`
	Skates    []Skate `json:"skates,omitempty"`
}

type Skate struct {
	Type      string `json:"type"` // Drop-In, Stick and Puck, etc.
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
}

// Return a fully initialized copy of the Rink
func NewRink(i int) (*Rink, error) {
	r := Rinks[i]
	err := r.GetSkates()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// We need a getter to setup the thing.
func (r *Rink) GetSkates() error {
	skates, found := Cache.Get(strconv.Itoa(r.RinkID))
	if found {
		r.Skates = skates.([]Skate)
		return nil
	}

	skates, err := fetchSkateData(r.API, r.Parser)
	if err != nil {
		log.Error(err)
	}
	Cache.Set(strconv.Itoa(r.RinkID), skates, cache.DefaultExpiration)

	return nil
}

func fetchSkateData(api string, parser string) ([]Skate, error) {
	var skates []Skate
	var p SkateParser

	resp, err := http.Get(api)
	if err != nil {
		log.Error("Could not retrieve data: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

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
