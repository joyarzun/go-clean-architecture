package presenter

import (
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type holidayPresenter struct{}

func New() usecases.HolidayPresenter {
	return &holidayPresenter{}
}

func (up *holidayPresenter) ResponseHoliday(holidays []*entities.Holiday) []*entities.Holiday {
	for _, u := range holidays {
		u.Name = cases.Title(language.Und).String(u.Name)
	}
	return holidays
}
