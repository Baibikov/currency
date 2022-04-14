package service

import (
	"context"

	"currency/internal/app/config"
	"currency/internal/app/repository"
	"currency/internal/app/types"
	"currency/pkg/currate_client"
)

type currency struct {
	apiConfig config.API
	storage *repository.Storage
}

func newCurrency(storage *repository.Storage, apiConfig config.API) *currency {
	return &currency{
		storage: storage,
		apiConfig: apiConfig,
	}
}

func (c *currency) Create(ctx context.Context, currency types.Currency) error {
	return c.storage.Currency.CreateCurrencyPair(ctx, currency)
}

func (c *currency) Update(ctx context.Context, currency types.Currency) error {
	return c.storage.Currency.UpdateCurrencyPair(ctx, currency)
}

func (c *currency) GetAll(ctx context.Context) ([]types.Currency, error) {
	return c.storage.Currency.GetCurrencyPairs(ctx)
}

func (c *currency) UpdateAll(ctx context.Context) error {
	currateClient := currate_client.New(currate_client.Config{
		Key: c.apiConfig.Key,
	})

	currencies, err := c.GetAll(ctx)
	if err != nil {
		return err
	}

	curPairs := make(currate_client.Pairs, 0, len(currencies))
	for _, cc := range currencies {
		curPairs = append(curPairs, currate_client.Pair{
			From: cc.From,
			To: cc.To,
		})
	}

	resp, err := currateClient.GetRates(curPairs)
	if err != nil {
		return err
	}

	for _, d := range resp.Data {
		err := c.Update(ctx, types.Currency{
			From: d.From,
			To: d.To,
			Well: d.Well,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
