package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func New() *Repository {
	return new(Repository)
}

func (r *Repository) WithPostgresqlDB(db *gorm.DB) *Repository {
	r.db = db
	return r
}
