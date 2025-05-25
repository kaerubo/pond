package postgres

import "database/sql"

type KeroRepository struct {
	db *sql.DB
}

func NewKeroRepository(db *sql.DB) *KeroRepository {
	return &KeroRepository{db: db}
}
