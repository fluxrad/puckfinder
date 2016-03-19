package api

import (
	"github.com/PuloV/ics-golang"
	log "github.com/Sirupsen/logrus"
)

type ICALParser struct {
	URL   string
	Start string
	End   string
}

func NewICALParser(url string) *ICALParser {
	p := &ICALParser{
		URL: url,
	}
	return p
}

func (i *ICALParser) Fetch() ([]byte, error) {
	// This is kind of janky. Should probably rework how we do fetch -> parse
	var nothing []byte

	log.Debug("ICAL parser. Not fetching data until parse is called.")
	return nothing, nil
}

func (i *ICALParser) Parse(b []byte) ([]*Skate, error) {
	var skates []*Skate

	parser := ics.New()
	parserChan := parser.GetInputChan()

	log.Debugf("fetching data from URL: %s", i.URL)
	parserChan <- i.URL

	outputChan := parser.GetOutputChan()
	go func() {
		for event := range outputChan {
			skates = append(skates, &Skate{
				Type:      event.GetSummary(),
				StartTime: event.GetStart(),
				EndTime:   event.GetEnd(),
			})
		}
	}()

	// wait to kill the main goroute
	parser.Wait()

	return skates, nil
}
