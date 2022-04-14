package pq

import (
	"context"

	"currency/internal/app/types"
)

func (p *PostgresRepository) CreateCurrencyPair(ctx context.Context, currency types.Currency) error {
	query := `
		insert into currency.pairs(currency_from, currency_to)
		values (:currency_from, :currency_to)
	`

	_, err := p.db.NamedExecContext(ctx, query, currency)
	return err
}

func (p *PostgresRepository) UpdateCurrencyPair(ctx context.Context, currency types.Currency) error {
	query := `
		update currency.pairs set well=:well,updated_at=now()
		where currency_from=:currency_from and currency_to=:currency_to
	`

	_, err := p.db.NamedExecContext(ctx, query, currency)
	return err
}

func (p *PostgresRepository) GetCurrencyPairs(ctx context.Context) (currencies []types.Currency, err error) {
	query := `
		select currency_from, currency_to, well, updated_at from currency.pairs
	`

	return currencies, p.db.SelectContext(ctx, &currencies, query)
}