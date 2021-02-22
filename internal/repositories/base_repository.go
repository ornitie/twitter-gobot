package repositories

import (
	"database/sql"
)

type GenericRepository interface {
	Create(query string) error
	Read(query string) (interface{}, error)
	Update(query string) (interface{}, error)
	Delete(query string) (interface{}, error)
}

type BaseRepository struct {
	db *sql.DB
}

func NewBaseRepository(db *sql.DB) *BaseRepository {
	return &BaseRepository{db}
}

func (repository BaseRepository) Create(query string) error {
	_, err := repository.db.Query(query)

	return err
}
