package api

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type SkateParser interface {
	Parse(b []byte) ([]*Skate, error)
}

type CalendarJSONParser struct{}

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
			StartTime: parseTime(v.Start),
			EndTime:   parseTime(v.End),
		}
		skates = append(skates, s)
	}

	return skates, nil
}

func parseTime(t string) int {
	return 0
}
