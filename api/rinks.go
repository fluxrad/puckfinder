package api

import (
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
	Type      string `json:"type"` // Drop-In, Stick and Puck, etc.
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
}

// Skates returns a list of Skates for a particular Rink. If the data doesn't
// currently exist in the cache, or has expired, it will be fetched, and
// cached.
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

// fetchSkateData fetches skate information from the target API and returns it
// as an slice of pointers to Skate
func fetchSkateData(api string, parser string) ([]*Skate, error) {
	var skates []*Skate
	var p SkateParser

	switch parser {
	case `calendarjson`:
		t := time.Now()
		p = NewCalendarJSONParser(api, t.Format("2006-01-02"), t.Add(720*time.Hour).Format("2006-01-02"))
	}

	data, err := p.Fetch()
	if err != nil {
		log.Error("Error fetching data: ", err)
	}

	skates, err = p.Parse(data)
	if err != nil {
		log.Error("Could not parse skates: %s", err)
	}

	return skates, nil
}
