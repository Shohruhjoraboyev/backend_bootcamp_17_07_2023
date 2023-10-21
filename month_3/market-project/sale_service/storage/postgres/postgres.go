package postgres

import (
	"context"
	"fmt"
	"sale_service/config"
	"sale_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db   *pgxpool.Pool
	sale *saleRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}

	return &strg{
		db: pool,
	}, nil
}

func (d *strg) Sale() storage.SaleI {
	if d.sale == nil {
		d.sale = NewSaleRepo(d.db)
	}
	return d.sale
}
