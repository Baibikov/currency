package service

import (
	"context"

	"currency/internal/app/config"
	"currency/internal/app/repository"
	"currency/internal/app/types"
)

type UseCase struct {
	Pair Pairer
}

type Pairer interface {
	Create(ctx context.Context, currency types.Currency) error
	Update(ctx context.Context, currency types.Currency) error
	GetAll(ctx context.Context) ([]types.Currency, error)
	UpdateAll(ctx context.Context) error
}


func New(storage *repository.Storage, apiConfig config.API) *UseCase {
	return &UseCase{
		Pair: newCurrency(storage, apiConfig),
	}
}