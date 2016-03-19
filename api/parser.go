package api

type SkateParser interface {
	Parse(b []byte) ([]*Skate, error)
	Fetch() ([]byte, error)
}
