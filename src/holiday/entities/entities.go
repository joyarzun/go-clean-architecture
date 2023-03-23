package entities

import "time"

type Holiday struct {
	Year int16     `json:"year"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
