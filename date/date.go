package date

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 0 {
		return errors.New("invalid Year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if 1 > month || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if 1 > day || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
