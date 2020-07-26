package geo

import (
	"errors"
	"unicode/utf8"
)

type Landmark struct {
	name string
	Coordinate
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if utf8.RuneCountInString(name) > 15 {
		return errors.New("name is too long")
	}
	l.name = name
	return nil
}
