package controller

import (
	"strings"
	"time"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse(time.RFC3339, strings.Trim(string(b), `"`))
	*t = Timestamp(ts)
	return err
}

type RequestHoliday struct {
	entities.Holiday
	Date Timestamp `json:"date"`
}
