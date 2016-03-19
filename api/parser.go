package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

type SkateParser interface {
	Parse(b []byte) ([]*Skate, error)
	Fetch() ([]byte, error)
}

type CalendarJSONParser struct {
	URL   string // URL to the API
	Start string // YYYY-MM-DD
	End   string // YYYY-MM-DD
}

type CalendarJSONEvents []*CalendarJSONEvent

type CalendarJSONEvent struct {
	AllDay                      string `json:"allDay"`
	End                         string `json:"end"`
	EndDate                     string `json:"endDate"`
	EndTime                     string `json:"endTime"`
	EventAnchorVenue            string `json:"eventAnchorVenue"`
	EventAnchorVenueDescription string `json:"eventAnchorVenueDescription"`
	EventCategory               string `json:"eventCategory"`
	EventClass                  string `json:"eventClass"`
	EventType                   string `json:"eventType"`
	ID                          string `json:"id"`
	Start                       string `json:"start"`
	StartDate                   string `json:"startDate"`
	StartTime                   string `json:"startTime"`
	Title                       string `json:"title"`
}

// NewCalendarJSONParser returns a pointer to an initialized
// CalendarJSONParser
func NewCalendarJSONParser(url string, start string, end string) *CalendarJSONParser {
	c := &CalendarJSONParser{
		URL:   url,
		Start: start,
		End:   end,
	}
	return c
}

func (c *CalendarJSONParser) Fetch() ([]byte, error) {
	query := fmt.Sprintf("%s&start=%s&end=%s", c.URL, c.Start, c.End)
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *CalendarJSONParser) Parse(b []byte) ([]*Skate, error) {
	var skates []*Skate
	var events CalendarJSONEvents

	if err := json.Unmarshal(b, &events); err != nil {
		log.Error("Could not parse json data: ", err)
		return nil, err
	}

	for _, v := range events {
		s := &Skate{
			Type:      v.Title,
			StartTime: parseTime(time.RFC3339, v.Start),
			EndTime:   parseTime(time.RFC3339, v.End),
		}
		skates = append(skates, s)
	}

	return skates, nil
}

func parseTime(format string, t string) int {
	tiem, err := time.Parse(format, t)
	if err != nil {
		log.Error("Couldn't parse time: ", err)
		return 0
	}

	u, err := strconv.Atoi(tiem.Format(time.UnixDate))
	if err != nil {
		log.Error("Couldn't convert time to unix time: ", err)
		return 0
	}

	return u
}
