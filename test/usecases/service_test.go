package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"gitlab.com/joyarzun/go-clean-architecture/test/mock"
	"gorm.io/gorm"
)

type holidayPresenter struct{}

func (up *holidayPresenter) ResponseHoliday(holidays []*entities.Holiday) []*entities.Holiday {
	return []*entities.Holiday{}
}

type dbMock struct {
	CreateWasCalled bool
}

func (db *dbMock) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return &gorm.DB{}
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
	var newDBMock dbMock
	var newRepository usecases.HolidayRepository
	var newPresenter usecases.HolidayPresenter
	var holidayService usecases.HolidayService

	BeforeEach(func() {
		newDBMock = dbMock{}
		newRepository = repository.New(&newDBMock)
		newPresenter = &holidayPresenter{}
		holidayService = usecases.New(&newRepository, &newPresenter)
	})

	It("should call DB Create when holiday create method is called", func() {
		Expect(holidayService).NotTo(BeNil())
		holidayService.Create(&mock.Holiday)
		Expect(newDBMock.CreateWasCalled).To(BeTrue())
	})

	//It("should find fist record seeking by id", func() {
	//	holidayService.Get(mock.Year)
	//	Expect()
	//})
})
