package usecases

import (
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
)

type HolidayPresenter interface {
	ResponseHoliday(h []*entities.Holiday) []*entities.Holiday
}
