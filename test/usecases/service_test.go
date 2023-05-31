package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"gitlab.com/joyarzun/go-clean-architecture/test/mock"
)

type holidayPresenter struct{}

func (up *holidayPresenter) ResponseHoliday(holidays []*entities.Holiday) []*entities.Holiday {
	return []*entities.Holiday{}
}

type dbMock struct {
	CreateWasCalled bool
	FindWasCalled   bool
}

func (db *dbMock) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	db.FindWasCalled = true
	return &gorm.DB{
		RowsAffected: 2,
	}
}

func (db *dbMock) Create(value interface{}) (tx *gorm.DB) {
	db.CreateWasCalled = true
	return &gorm.DB{
		RowsAffected: 1,
	}
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

	It("should create a new holiday", func() {
		result := newDBMock.Create([]entities.Holiday{})
		Expect(result.RowsAffected).To(Equal(int64(1)))
	})

	It("should call DB Create when holiday create method is called", func() {
		Expect(holidayService).NotTo(BeNil())
		holidayService.Create(&mock.Holiday)
		Expect(newDBMock.CreateWasCalled).To(BeTrue())
	})

	It("should find all by year", func() {
		result := newDBMock.Find([]entities.Holiday{}, "year = 2023")
		Expect(result.RowsAffected).To(Equal(int64(2)))
	})

	It("should find all records seeking by year", func() {
		holidayService.Get(mock.Year)
		Expect(newDBMock.FindWasCalled).To(BeTrue())
	})

})
