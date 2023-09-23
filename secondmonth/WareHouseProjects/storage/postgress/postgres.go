package postgres

import (
	"WareHouseProjects/config"
	"WareHouseProjects/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db       *pgxpool.Pool
	branches *branchRepo
	category *categoryRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnections

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (b *store) Branch() storage.BranchesI {
	if b.branches == nil {
		b.branches = NewBranchRepo(b.db)
	}
	return b.branches
}

func (b *store) Category() storage.CategoriesI {
	if b.category == nil {
		b.category = NewCategoryRepo(b.db)
	}
	return b.category
}

func (s *store) Close() {
	s.db.Close()
}
