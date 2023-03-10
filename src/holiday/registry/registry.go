package registry

import (
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/controller"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/presenter"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Holiday: r.NewHolidayController(),
	}
}

func (r *registry) NewHolidayController() controller.HolidayController {
	newRepository := repository.New(r.db)
	newpresenter := presenter.New()
	holidayService := usecases.New(&newRepository, &newpresenter)

	return controller.New(holidayService)
}
