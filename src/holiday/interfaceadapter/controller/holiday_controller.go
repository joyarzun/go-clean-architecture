package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
)

type holidayController struct {
	holidayService usecases.HolidayService
}

type HolidayController interface {
	GetHolidays(c echo.Context) error
	CreateHoliday(c echo.Context) error
}

func New(us usecases.HolidayService) HolidayController {
	return &holidayController{us}
}

func (hc *holidayController) GetHolidays(c echo.Context) error {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		return err
	}

	holidays, err := hc.holidayService.Get(int16(year))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, holidays)
}

func (hc *holidayController) CreateHoliday(c echo.Context) error {

	requestHoliday := new(RequestHoliday)
	if err := c.Bind(requestHoliday); err != nil {
		return err
	}

	holiday, err := hc.holidayService.Create(requestHoliday.ToHoliday())

	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, holiday)
}
