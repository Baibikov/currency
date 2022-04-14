package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"currency/internal/app/repository/pq"
	"currency/internal/app/types"
)

type Storage struct {
	Currency Currency
}

type Currency interface {
	CreateCurrencyPair(ctx context.Context, currency types.Currency) error
	UpdateCurrencyPair(ctx context.Context, currency types.Currency) error
	GetCurrencyPairs(ctx context.Context) (currencies []types.Currency, err error)
}

func New(db *sqlx.DB) *Storage{
	return &Storage{
		Currency: pq.New(db),
	}
}
