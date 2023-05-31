package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/test/mock"
)

type dbMock struct{}

func (db *dbMock) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return &gorm.DB{}
}

func (db *dbMock) Create(value interface{}) (tx *gorm.DB) {
	return &gorm.DB{}
}

var _ = Describe("Repository", func() {

	It("Find all by year", func() {
		db, err := gorm.Open(sqlite.Open("mock.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})


		holidayRepositoryMock := repository.New(db)
		holidayMock, err := holidayRepositoryMock.Create(&mock.Holiday)

		response, err := holidayRepositoryMock.FindAllByYear(int16(2023))
		
		db.Delete(entities.Holiday{}, 2023)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(response[0]).To(Equal(holidayMock))
	})
})
