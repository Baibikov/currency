package pq

import "github.com/jmoiron/sqlx"

type PostgresRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}