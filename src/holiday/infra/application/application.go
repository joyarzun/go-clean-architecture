package application

import (
	"fmt"

	"github.com/labstack/echo"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/datastore"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/router"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/registry"
)

func Run() {
	port := "3000"

	reg := registry.NewRegistry(datastore.NewDB())
	server := router.NewRouter(echo.New(), reg.NewAppController())
	server.Logger.Fatal(server.Start(":" + port))
	fmt.Println("Server listen at http://localhost" + ":" + port)
}
