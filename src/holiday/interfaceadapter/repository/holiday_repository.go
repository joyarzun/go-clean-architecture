package repository

import (
	"time"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"gorm.io/gorm"
)

type holidayRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) usecases.HolidayRepository {
	return &holidayRepository{db}
}

type Holiday struct {
	entities.Holiday
	Date string `json:"date"`
}

func (hr *holidayRepository) FindAllByYear(year int16) ([]*entities.Holiday, error) {
	var dbHoliday Holiday
	err := hr.db.First(&dbHoliday, "year = ?", year).Error

	if err != nil {
		return nil, err
	}

	parsedDate, err := time.Parse("2006-01-02 15:04:05+00:00", dbHoliday.Date)

	if err != nil {
		return nil, err
	}

	holiday := entities.Holiday{
		Name: dbHoliday.Name,
		Year: dbHoliday.Year,
		Date: parsedDate,
	}

	holidays := []*entities.Holiday{&holiday}

	return holidays, nil
}

func (hr *holidayRepository) Create(u *entities.Holiday) (*entities.Holiday, error) {
	err := hr.db.Create(u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
