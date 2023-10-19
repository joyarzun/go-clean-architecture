package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/infra/datastore"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/usecases"
	"gitlab.com/joyarzun/go-clean-architecture/test/mock"
)

var db datastore.Storei
var holidayRepositoryMock usecases.HolidayRepository
var holidayMock *entities.Holiday

var _ = Describe("Repository", Ordered, func() {

	BeforeAll(func() {
		db = datastore.New("")
		err := db.Error()

		if err != nil {
			panic(err)
		}

		holidayRepositoryMock = repository.New(db)
	})

	When("There is at least one holiday registered on database", func() {
		BeforeEach(func() {
			var err error
			holidayMock, err = holidayRepositoryMock.Create(&mock.Holiday)
			if err != nil {
				panic(err)
			}
		})

		It("Find all by year", func() {
			response, err := holidayRepositoryMock.FindAllByYear(int16(2023))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(response[0]).To(BeEquivalentTo(holidayMock))
		})
	})

	When("There is no holiday registered on database", func() {
		It("Find all by year", func() {
			response, _ := holidayRepositoryMock.FindAllByYear(int16(2023))
			Expect(response).To(BeNil())
		})
	})

	AfterEach(func() {
		db.Delete(entities.Holiday{}, int16(2023))
	})

})
