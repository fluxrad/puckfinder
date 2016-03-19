package api

import (
	"time"

	log "github.com/Sirupsen/logrus"
	cache "github.com/patrickmn/go-cache"
)

var Rinks map[int]*Rink
var Cache *cache.Cache

func init() {
	log.SetLevel(log.DebugLevel)

	Rinks = make(map[int]*Rink)
	Cache = cache.New(10*time.Minute, 30*time.Second)

	// Fetch this from a DB later (bolt?)
	Rinks[1] = &Rink{
		RinkID:    1,
		ShortName: "University of Denver",
		URL:       "http://recreation.du.edu",
		API:       "http://denveruniv-web.ungerboeck.com/Calendar/Default.aspx?&EventClassFilter=classFilter1&EventFilter=&EventSearchTerms=&EventFormat=FULLCALENDARJSON&EventSingleClassFilter=IAH",
		Parser:    "calendarjson",
		Timeout:   600,
	}

	Rinks[2] = &Rink{
		RinkID:    2,
		ShortName: "Boulder Valley Ice",
		URL:       "http://www.bvice.com",
		API:       "http://bvice.pucksystems.com/ical_feed?tags=35499%252C492957%252C37560%252C37564%252C37561%252C37562%252C41822%252C101399%252C333682%252C35528%252C35529%252C35530",
		Parser:    "ical",
		Timeout:   600,
	}
}
