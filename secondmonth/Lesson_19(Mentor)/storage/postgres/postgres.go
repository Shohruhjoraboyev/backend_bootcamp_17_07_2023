package postgres

import (
	"context"
	"fmt"
	"projectWithGinAndSwagger/config"
	"projectWithGinAndSwagger/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db          *pgxpool.Pool
	branches    *branchRepo
	staffes     *staffRepo
	sales       *salesRepo
	transaction *transactionRepo
	staffTarifs *staffTarifRepo
}

// BiznesLogic implements storage.StorageI.
func (*store) BiznesLogic() storage.BiznesLogicI {
	panic("unimplemented")
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

func (b *store) Staff() storage.StaffesI {
	if b.staffes == nil {
		b.staffes = NewStaffRepo(b.db)
	}
	return b.staffes
}

func (s *store) Sales() storage.SalesI {
	if s.sales == nil {
		s.sales = NewSaleRepo(s.db)
	}
	return s.sales
}

func (b *store) Transaction() storage.TransactionI {
	if b.transaction == nil {
		b.transaction = NewTransactionRepo(b.db)
	}
	return b.transaction
}

func (b *store) StaffTarif() storage.StaffTarifsI {
	if b.staffTarifs == nil {
		b.staffTarifs = NewStaffTarifRepo(b.db)
	}
	return b.staffTarifs
}

func (s *store) Close() {
	s.db.Close()
}
