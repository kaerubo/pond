package postgres

import (
	"database/sql"
)

type keroRepository struct {
	db *sql.DB
}

func NewKeroRepository(db *sql.DB) *keroRepository {
	return &keroRepository{db: db}
}
