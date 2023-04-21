package main_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type holidayPresenter struct{}

func (up *holidayPresenter) ResponseHoliday(holidays []*entities.Holiday) []*entities.Holiday {
	for _, u := range holidays {
		u.Name = cases.Title(language.Und).String(u.Name)
	}
	return holidays
}

func newHolidayPresenterMock() usecases.HolidayPresenter {
	return &holidayPresenter{}
}

type dbMock struct {
	CreateWasCalled bool
}

func (db *dbMock) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return &gorm.DB{}
}

func (db *dbMock) Create(value interface{}) (tx *gorm.DB) {
	db.CreateWasCalled = true
	dbi := gorm.DB{}
	return &dbi
}

var _ = Describe("Usecase Service", func() {
	It("should call DB Create when holiday create method is called", func() {
		newdbMock := dbMock{}
		newRepository := repository.New(&newdbMock)
		newpresenter := newHolidayPresenterMock()
		holidayService := usecases.New(&newRepository, &newpresenter)
		Expect(holidayService).NotTo(BeNil())
		holiday := entities.Holiday{
			Year: 2023,
			Name: "a",
			Date: time.Now(),
		}
		holidayService.Create(&holiday)
		Expect(newdbMock.CreateWasCalled).To(BeTrue())
	})
})
