package repository

import (
	"time"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"gorm.io/gorm"
)

type holidayRepository struct {
	db GormDB
}

func New(db GormDB) usecases.HolidayRepository {
	return &holidayRepository{db: db}
}

type Holiday struct {
	entities.Holiday
	Date string `json:"date"`
}

func (hr *holidayRepository) FindAllByYear(year int16) ([]*entities.Holiday, error) {
	var dbHoliday []Holiday
	var parsedDate time.Time
	var holidays []*entities.Holiday

	result := hr.db.Find(&dbHoliday, "year = ?", year)
	err := result.Error

	if err != nil {
		return nil, err
	}

	for _, dbh := range dbHoliday {
		parsedDate, err = time.Parse("2006-01-02 15:04:05+00:00", dbh.Date)
		if err == nil {
			holidays = append(holidays, &entities.Holiday{
				Year: dbh.Year,
				Name: dbh.Name,
				Date: parsedDate,
			})
		}
	}
	return holidays, nil
}

func (hr *holidayRepository) Create(u *entities.Holiday) (*entities.Holiday, error) {
	err := hr.db.Create(u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

type GormDB interface {
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
}
