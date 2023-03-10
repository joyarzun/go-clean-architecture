package controller

type AppController struct {
	Holiday interface{ HolidayController }
}
