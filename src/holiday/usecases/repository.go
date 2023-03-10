package usecases

import "gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"

type HolidayRepository interface {
	FindAllByYear(year int16) ([]*entities.Holiday, error)
	Create(u *entities.Holiday) (*entities.Holiday, error)
}
