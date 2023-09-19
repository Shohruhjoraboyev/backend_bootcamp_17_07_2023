package postgres

import (
	"context"
	"fmt"
	"projectWithGinAndSwagger/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type transactionRepo struct {
	db *pgxpool.Pool
}

// GetTopStaffs implements storage.TransactionI.
func (*transactionRepo) GetTopStaffs(models.TopWorkerRequest) (map[string]models.StaffTop, error) {
	panic("unimplemented")
}

func NewTransactionRepo(db *pgxpool.Pool) *transactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (t *transactionRepo) CreateTransaction(req *models.CreateTransaction) (string, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO transactions(id,type,amount,source_type,text,sale_id,staff_id,created_at)
	VALUES($1, $2, $3, $4, $5, $6,$7,now())
	`
	_, err := t.db.Exec(context.Background(), query,
		id,
		req.Type,
		req.Amount,
		req.Source_type,
		req.Text,
		req.Sale_id,
		req.Staff_id,
	)
	if err != nil {
		return "Error Create Transaction", err
	}

	return id, nil
}

func (t *transactionRepo) GetTransaction(req *models.IdRequest) (resp *models.Transaction, err error) {
	query := `SELECT *FROM transactions WHERE "id" = $1`
	transaction := models.Transaction{}
	err = t.db.QueryRow(context.Background(), query, req.Id).Scan(
		&transaction.Id,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Source_type,
		&transaction.Text,
		&transaction.Sale_id,
		&transaction.Staff_id,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("transaction not found")

	}
	return &transaction, err
}

func (t *transactionRepo) GetAllTransaction(req *models.GetAllTransactionRequest) (*models.GetAllTransactionResponse, error) {
	var response models.GetAllTransactionResponse
	response.Transaction = make([]models.Transaction, 0)

	query := `
		SELECT id, type, amount, source_type, text, sale_id, staff_id, created_at, updated_at
		FROM transactions
		WHERE text ILIKE '%' || $1 || '%'
		LIMIT $2 OFFSET $3
	`

	rows, err := t.db.Query(context.Background(), query, req.Text, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Source_type,
			&transaction.Text,
			&transaction.Sale_id,
			&transaction.Staff_id,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}
		response.Transaction = append(response.Transaction, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over transactions: %w", err)
	}

	countQuery := `
		SELECT COUNT(*)
		FROM transactions
		WHERE text ILIKE '%' || $1 || '%'
	`
	err = t.db.QueryRow(context.Background(), countQuery, req.Text).Scan(&response.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction count: %w", err)
	}

	return &response, nil
}

func (t *transactionRepo) UpdateTransaction(req *models.Transaction) (string, error) {
	query := `UPDATE transactions SET amount=$1,source_type=$2,text=$3,sale_id=$4,staff_id=$5,updated_at=now(), WHERE id=$6 RETURNING id`
	result, err := t.db.Exec(context.Background(), query, req.Amount, req.Source_type, req.Text, req.Sale_id, req.Staff_id, req.Id)
	if err != nil {
		return "Error update transaction", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Transaction not found")
	}

	return req.Id, nil
}

func (t *transactionRepo) DeleteTransaction(req *models.IdRequest) (string, error) {
	query := `DELETE FROM transactions WHERE id=$1 RETURNING id`

	result, err := t.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "Error delete transaction", err
	}
	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch not found")
	}

	return req.Id, nil

}
