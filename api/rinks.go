package api

import (
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/patrickmn/go-cache"
)

// Store this in a database eventually
type Rink struct {
	RinkID    int    `json:"id"`
	ShortName string `json:"shortName"`
	URL       string `json:"url"`
	API       string
	Parser    string
	Timeout   int
}

type RinkData struct {
	*Rink  `json:"rink"`
	Skates []*Skate `json:"skates"`
}

type Skate struct {
	Type      string `json:"type"` // Drop-In, Stick and Puck, etc.
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
}

func fetchRinkInfo(i int) (*RinkData, error) {
	var d RinkData

	// SELECT rinkID, api, parser, timeout FROM rink_info
	d.Rink = Rinks[i]

	skates, found := Cache.Get(strconv.Itoa(d.RinkID))
	if found {
		d.Skates = skates.([]*Skate)
		return &d, nil
	}

	skates, err := fetchSkates(&d)
	if err != nil {
		log.Error(err)
	}
	Cache.Set(strconv.Itoa(d.RinkID), skates, cache.DefaultExpiration)

	return &d, nil
}

func fetchSkates(r *RinkData) ([]*Skate, error) {
	return nil, nil
}
