package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/presenter"
	"gitlab.com/joyarzun/go-clean-architecture/test/mock"
)

var _ = Describe("Holiday Presenter", func() {
	It("Should format the list of holidays", func() {
		holidayPresenterMock := presenter.New()
		holidayFormatted := holidayPresenterMock.ResponseHoliday(mock.Holidays)
		Expect(&holidayFormatted).To(Equal(&mock.HolidaysExpected))
	})
})
