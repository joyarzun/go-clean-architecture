package mock

import (
	"time"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
)

var Year = int16(2023)

var Holiday = entities.Holiday{
	Year: 2023,
	Name: "a",
	Date: time.Date(2023, time.January, 01, 00, 00, 00, 00, time.UTC),
}

var Holidays = []*entities.Holiday{
	{
		Year: 2023,
		Name: "new year",
		Date: time.Date(2023, time.January, 01, 00, 00, 00, 00, time.UTC),
	},
	{
		Year: 2023,
		Name: "good friday",
		Date: time.Date(2023, time.April, 07, 00, 00, 00, 00, time.UTC),
	},
}

var HolidaysExpected = []*entities.Holiday{
	{
		Year: 2023,
		Name: "New Year",
		Date: time.Date(2023, time.January, 01, 00, 00, 00, 00, time.UTC),
	},
	{
		Year: 2023,
		Name: "Good Friday",
		Date: time.Date(2023, time.April, 07, 00, 00, 00, 00, time.UTC),
	},
}
