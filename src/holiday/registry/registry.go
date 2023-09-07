package registry

import (
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/datastore"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/controller"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/presenter"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
)

type registry struct {
	db datastore.Storei
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db datastore.Storei) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Holiday: r.NewHolidayController(),
	}
}

func (r *registry) NewHolidayController() controller.HolidayController {
	newRepository := repository.New(r.db)
	newPresenter := presenter.New()
	holidayService := usecases.New(&newRepository, &newPresenter)

	return controller.New(holidayService)
}
