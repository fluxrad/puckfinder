package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

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
func NewCalendarJSONParser(url string) *CalendarJSONParser {
	t := time.Now()
	start := t.Format("2006-01-02")
	end := t.Add(720 * time.Hour).Format("2006-01-02")

	c := &CalendarJSONParser{
		URL:   url,
		Start: start,
		End:   end,
	}
	return c
}

// Fetch fetches data from the orgiin server and returns a byte
// slice containing the body of the request.
func (c *CalendarJSONParser) Fetch() ([]byte, error) {
	query := fmt.Sprintf("%s&start=%s&end=%s", c.URL, c.Start, c.End)
	log.Debugf("querying url: %s", query)

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

	const ubFormat = "2006-01-02T15:04:05" // UngerBoeck whatever

	for _, v := range events {
		s := &Skate{
			Type:      v.Title,
			StartTime: parseTime(ubFormat, v.Start),
			EndTime:   parseTime(ubFormat, v.End),
		}
		skates = append(skates, s)
	}

	return skates, nil
}

func parseTime(format string, t string) time.Time {
	tiem, err := time.Parse(format, t)
	if err != nil {
		log.Error("Couldn't parse time: ", err)
		return time.Now()
	}

	return tiem
}
