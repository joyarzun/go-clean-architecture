package repository

type Repository interface {
	Find(dest interface{}, conds ...interface{})
	Create(value interface{})
	Error() error
}
