package geo

import "errors"

type Coordinate struct {
	latitude  float64
	longitude float64
}

// Getters
func (c *Coordinate) Latitude() float64 {
	return c.latitude
}
func (c *Coordinate) Longitude() float64 {
	return c.longitude
}

// Setters
func (c *Coordinate) SetLatitude(latitude float64) error {
	if -90 > latitude || latitude > 90 {
		return errors.New("latitude must -90 ~ 90")
	}
	c.latitude = latitude
	return nil
}
func (c *Coordinate) SetLongitude(longitude float64) error {
	if -180 > longitude || longitude > 180 {
		return errors.New("longitude must -180 ~ 180")
	}
	c.longitude = longitude
	return nil
}
