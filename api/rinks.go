package api

import (
	"errors"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/patrickmn/go-cache"
)

// Rink is an object containing information about a particular facility
type Rink struct {
	RinkID    int    `json:"id"`
	ShortName string `json:"shortName"`
	URL       string `json:"url"`
	API       string `json:"api"`
	Parser    string `json:"parser"`
	Timeout   int    `json:"timeout"`
}

// Skate contains information about a specific skating event
type Skate struct {
	Type      string    `json:"type"` // Drop-In, Stick and Puck, etc.
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// Skates returns a list of Skates for a particular Rink. If the data doesn't
// currently exist in the cache, or has expired, it will be fetched, and
// cached.
func (r *Rink) Skates() []*Skate {
	skates, found := Cache.Get(strconv.Itoa(r.RinkID))
	if found {
		log.Debug("Found cached data. Returning it")
		return skates.([]*Skate)
	}

	log.Debug("Did not find cached data. Fetching")
	skates, err := fetchSkateData(r.API, r.Parser)
	if err != nil {
		log.Error(err)
	}
	Cache.Set(strconv.Itoa(r.RinkID), skates, cache.DefaultExpiration)

	return skates.([]*Skate)
}

// fetchSkateData fetches skate information from the target API and returns it
// as an slice of pointers to Skate
func fetchSkateData(api string, parser string) ([]*Skate, error) {
	var skates []*Skate
	var p SkateParser

	switch parser {
	case `calendarjson`:
		p = NewCalendarJSONParser(api)
	case `ical`:
		p = NewICALParser(api)
	default:
		err := errors.New("No such API type")
		log.Error(err)
		return nil, err
	}

	skateData, err := p.Fetch()
	if err != nil {
		log.Error("Error fetching data: ", err)
	}

	skates, err = p.Parse(skateData)
	if err != nil {
		log.Error("Could not parse skates: %s", err)
	}

	return skates, nil
}
