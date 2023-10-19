package datastore

import (
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/entities"
	"gitlab.com/joyarzun/go-clean-architecture/src/holiday/interfaceadapter/repository"
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
	DB  map[string]*repository.Holiday
	err error
}

func New(dbfile string) Storei {
	DB := make(map[string]*repository.Holiday)
	return &Store{DB: DB}
}

func (s *Store) Find(dest interface{}, conds ...interface{}) {
	year := conds[1].(int16)
	result := make([]repository.Holiday, 0)
	for _, v := range s.DB {
		if v.Year == year {
			result = append(result, *v)
		}
	}
	*dest.(*[]repository.Holiday) = result
}

func (s *Store) Create(value interface{}) {
	s.DB[value.(*entities.Holiday).Name] = EntityToRepository(value.(*entities.Holiday))
}

func (s *Store) Error() error {
	return s.err
}

func (s *Store) Exec(sql string) {

}

func (s *Store) Delete(value interface{}, conds ...interface{}) {
	year := conds[0].(int16)
	for _, v := range s.DB {
		if v.Year == year {
			delete(s.DB, v.Name)
		}
	}
}

func EntityToRepository(e *entities.Holiday) *repository.Holiday {
	return &repository.Holiday{
		Holiday: *e, Date: e.Date.Format("2006-01-02 15:04:05+00:00"),
	}
}
