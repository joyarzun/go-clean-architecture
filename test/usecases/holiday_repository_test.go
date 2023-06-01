package main

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/test/mock"
)

var _ = Describe("Repository", func() {

	It("Find all by year", func() {
		db, err := gorm.Open(sqlite.Open("mock.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		_ = db.Exec("CREATE TABLE `holidays` (`id` INTEGER PRIMARY KEY,`year` INTEGER NOT NULL,`name` TEXT NOT NULL,`date` TEXT NOT NULL)")
		if err != nil {
			log.Fatal(err)
		}

		holidayRepositoryMock := repository.New(db)
		holidayMock, _ := holidayRepositoryMock.Create(&mock.Holiday)
		response, err := holidayRepositoryMock.FindAllByYear(int16(2023))

		db.Delete(entities.Holiday{}, 2023)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(response[0]).To(BeEquivalentTo(holidayMock))
	})
})
