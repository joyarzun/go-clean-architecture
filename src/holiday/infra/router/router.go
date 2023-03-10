package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/holiday/:year", func(context echo.Context) error { return c.Holiday.GetHolidays(context) })
	e.POST("/holiday", func(context echo.Context) error { return c.Holiday.CreateHoliday(context) })

	return e
}
