package datastore

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DBFILE = "gorm.db"

type Storei interface {
	Find(dest interface{}, conds ...interface{})
	Create(value interface{})
	Error() error
	Exec(sql string)
	Delete(value interface{}, conds ...interface{})
}

type Store struct {
	DB *gorm.DB
}

func New(dbfile string) Storei {
	DB, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}

	return &Store{DB}
}

func (s *Store) Find(dest interface{}, conds ...interface{}) {
	s.DB.Find(dest, conds)
}

func (s *Store) Create(value interface{}) {
	s.DB.Create(value)
}

func (s *Store) Error() error {
	return s.DB.Error
}

func (s *Store) Exec(sql string) {
	s.DB.Exec(sql)
}

func (s *Store) Delete(value interface{}, conds ...interface{}) {
	s.DB.Delete(value, conds)
}
