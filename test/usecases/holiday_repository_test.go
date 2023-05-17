package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//var db *sql.DB

type dbMock struct{}

func (db *dbMock) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return &gorm.DB{}
}

func (db *dbMock) Create(value interface{}) (tx *gorm.DB) {
	return &gorm.DB{}
}

var _ = Describe("Repository", func() {

	It("Find all by year", func() {
		db, mock, err := sqlmock.New()
		sqlite.Dialector{
			DriverName: "",
			DSN:        "",
			Conn:       nil,
		}
		dialector := sqlite.Dialector{DSN: "gorm.db", Conn: sqlx.DB{DB: db}, sk}
		gdb, err := gorm.Open(dialector, &gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())
		defer db.Close()

		holidayRepository := repository.HolidayRepository{DB: gdb}

		mock.ExpectBegin()

		mock.ExpectExec("UPDATE holidays").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO holidays").
			WithArgs("2023", "New Year", "2023-01-01T00:00:00Z").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rows := sqlmock.NewRows([]string{"year", "name", "date"}).
			AddRow("2023", "New Year", "2023-01-01T00:00:00Z")

		mock.ExpectQuery("^SELECT * FROM holidays").WillReturnRows(rows)
		mock.ExpectCommit()

		response, err := holidayRepository.FindAllByYear(int16(2023))

		Expect(err).ShouldNot(HaveOccurred())
		Expect(response).Should(Not(BeEmpty()))
	})
})
