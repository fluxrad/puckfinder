package api

type SkateParser interface {
	Parse([]byte) ([]Skate, error)
}

type CalendarJSON struct {
}

func (c *CalendarJSON) Parse([]byte) ([]Skate, error) {
	return nil, nil
}
