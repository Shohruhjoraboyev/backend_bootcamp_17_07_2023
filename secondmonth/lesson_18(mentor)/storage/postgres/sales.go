package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type salesRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *salesRepo {
	return &salesRepo{
		db: db,
	}
}

func (s *salesRepo) CreateSale(req *models.CreateSale) (string, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO sales(id,client_name,branch_id,shop_assistant_id,cashier_id,price,payment_type,status,created_at)
	VALUES($1, $2, $3, $4, $5, $6,$7,$8,now())
	`

	if req.Shop_asisstant_id == "" {
		_, err := s.db.Exec(context.Background(), query,
			id,
			req.ClientName,
			req.Branch_id,
			nil,
			req.Cashier_id,
			req.Price,
			req.Payment_Type,
			req.Status,
		)

		if err != nil {
			return "Error Create Sale", err
		}
	} else if req.Cashier_id == "" {
		_, err := s.db.Exec(context.Background(), query,
			id,
			req.Branch_id,
			req.Shop_asisstant_id,
			nil,
			req.Price,
			req.Payment_Type,
			req.Status,
		)

		if err != nil {
			return "Error Create Sale", err
		}
	} else {
		_, err := s.db.Exec(context.Background(), query,
			id,
			req.Branch_id,
			req.Shop_asisstant_id,
			req.Cashier_id,
			req.Price,
			req.Payment_Type,
			req.Status,
		)

		if err != nil {
			return "Error Create Sale", err
		}
	}

	return id, nil
}
func (s *salesRepo) GetSale(req *models.IdRequest) (resp *models.Sales, err error) {
	query := `SELECT * FROM sales WHERE id=$1`
	sale := models.Sales{}
	err = s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&sale.Id,
		&sale.ClientName,
		&sale.Branch_id,
		&sale.Shop_asisstent_id,
		&sale.Cashier_id,
		&sale.Price,
		&sale.Payment_Type,
		&sale.Status,
		&sale.CreatedAt,
		&sale.Updated_at,
	)

	if err != nil {
		return nil, fmt.Errorf("Sale not found")
	}
	return &sale, nil
}

func (c *salesRepo) GetAllSale(req *models.GetAllSaleRequest) (*models.GetAllSaleResponse, error) {
	params := make(map[string]interface{})
	filter := ""

	var cashier_id sql.NullString
	var shop_assistant_id sql.NullString

	selectQuery := `
    SELECT *  FROM "sales"
    `

	if req.Client_name != "" {
		filter += ` WHERE "client_name" ILIKE '%' || :search || '%' `
		params["search"] = req.Client_name
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	offset := (req.Page - 1) * limit

	params["limit"] = limit
	params["offset"] = offset

	query := selectQuery + filter + " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := c.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp := &models.GetAllSaleResponse{}
	count := 0
	for rows.Next() {
		var sale models.Sales
		count++
		err := rows.Scan(
			&sale.Id,
			&sale.ClientName,
			&sale.Branch_id,
			&shop_assistant_id,
			&cashier_id,
			&sale.Price,
			&sale.Payment_Type,
			&sale.Status,
			&sale.CreatedAt,
			&sale.Updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Check if the fields are null and set them to empty strings if necessary
		if cashier_id.Valid {
			sale.Cashier_id = cashier_id.String
		}
		if shop_assistant_id.Valid {
			sale.Shop_asisstent_id = shop_assistant_id.String
		}

		resp.Sales = append(resp.Sales, sale)
	}

	resp.Count = count
	return resp, nil
}

func (s *salesRepo) UpdateSale(req *models.Sales) (string, error) {
	// edit $1, there is no name
	query := `UPDATE sales SET name=$1,client_name=$2,branch_id=$3,shop_assistent_id=$4,cashier_id=$5,price=$6,payment_type=$7,status=$8,updateed_at=now() WHERE id=$9 RETURING Id`
	result, err := s.db.Exec(context.Background(), query, req.ClientName, req.Branch_id, req.Shop_asisstent_id, req.Cashier_id, req.Price, req.Payment_Type, req.Status, req.Id)
	if err != nil {
		return "Error Update Sale", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Sale not found")
	}
	return req.Id, nil
}

func (s *salesRepo) DeleteSale(req *models.IdRequest) (string, error) {
	query := `DELETE FROM sales WHERE id=$1 RETURNING id`
	result, err := s.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "Error delete sale", err
	}
	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Sale not found")
	}
	return req.Id, nil
}
