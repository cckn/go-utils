package calendar

import (
	"fmt"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	titleLen := utf8.RuneCountInString(title)
	if titleLen > 20 {
		return fmt.Errorf("title is to long : %d char", titleLen)
	}
	e.title = title
	return nil
}
