package repository

import (
	"time"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
)

type HolidayRepository struct {
	DB Repository
}

func New(db Repository) usecases.HolidayRepository {
	return &HolidayRepository{DB: db}
}

type Holiday struct {
	entities.Holiday
	Date string `json:"date"`
}

func (hr *HolidayRepository) FindAllByYear(year int16) ([]*entities.Holiday, error) {
	var dbHoliday []Holiday
	var parsedDate time.Time
	var holidays []*entities.Holiday

	hr.DB.Find(&dbHoliday, "year = ?", year)
	err := hr.DB.Error()

	if err != nil {
		return nil, err
	}

	for _, dbh := range dbHoliday {
		parsedDate, err = time.Parse("2006-01-02 15:04:05+00:00", dbh.Date)

		if err != nil {
			return nil, err
		}

		holidays = append(holidays, &entities.Holiday{
			Year: dbh.Year,
			Name: dbh.Name,
			Date: parsedDate,
		})
	}
	return holidays, nil
}

func (hr *HolidayRepository) Create(u *entities.Holiday) (*entities.Holiday, error) {
	hr.DB.Create(u)
	err := hr.DB.Error()

	if err != nil {
		return nil, err
	}

	return u, nil
}
