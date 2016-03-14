package api

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var Rinks map[int]*Rink
var Cache *cache.Cache

func init() {
	Rinks = make(map[int]*Rink)
	Cache = cache.New(10*time.Minute, 30*time.Second)

	// Fetch this from a DB later (bolt?)
	Rinks[1] = &Rink{
		RinkID:    1,
		ShortName: "University of Denver",
		URL:       "http://recreation.du.edu",
		API:       "http://denveruniv-web.ungerboeck.com/Calendar/Default.aspx?EventClassFilter=classFilter1&EventFormat=FULLCALENDARJSON",
		Parser:    "calendarjson",
		Timeout:   600,
	}
}
