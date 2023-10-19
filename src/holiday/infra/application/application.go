package application

import (
	"fmt"

	"github.com/labstack/echo"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/datastore"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/router"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/controller"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/presenter"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
)

func Run() {
	port := "3000"

	holidayRepository := repository.New(datastore.New(datastore.DBFILE))
	holidayPresenter := presenter.New()
	holidayService := usecases.New(&holidayRepository, &holidayPresenter)

	mainController := controller.AppController{
		Holiday: controller.New(holidayService),
	}

	server := router.NewRouter(echo.New(), mainController)
	server.Logger.Fatal(server.Start(":" + port))
	fmt.Println("Server listen at http://localhost" + ":" + port)
}
