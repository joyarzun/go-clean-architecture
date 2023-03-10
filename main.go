package main

import (
	"fmt"

	"github.com/labstack/echo"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/datastore"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/router"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/registry"
)

func main() {

	db := datastore.NewDB()
	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	port := "3000"
	fmt.Println("Server listen at http://localhost" + ":" + port)
	e.Logger.Fatal(e.Start(":" + port))
}
