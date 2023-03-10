package usecases

import "gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"

type holidayService struct {
	HolidayRepository HolidayRepository
	HolidayPresenter  HolidayPresenter
}

type HolidayService interface {
	Get(year int16) ([]*entities.Holiday, error)
	Create(holiday *entities.Holiday) (*entities.Holiday, error)
}

func New(holidayRepository *HolidayRepository, holidayPresenter *HolidayPresenter) HolidayService {
	return &holidayService{
		HolidayRepository: *holidayRepository,
		HolidayPresenter:  *holidayPresenter,
	}
}

func (us *holidayService) Get(year int16) ([]*entities.Holiday, error) {
	holidays, err := us.HolidayRepository.FindAllByYear(year)
	if err != nil {
		return nil, err
	}

	return us.HolidayPresenter.ResponseHoliday(holidays), nil
}

func (us *holidayService) Create(holiday *entities.Holiday) (*entities.Holiday, error) {
	holiday, err := us.HolidayRepository.Create(holiday)
	if err != nil {
		return nil, err
	}
	return holiday, nil
}
