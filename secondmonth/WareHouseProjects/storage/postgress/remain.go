package postgres

import (
	"WareHouseProjects/models"
	"WareHouseProjects/pkg/helper"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type remainRepo struct {
	db *pgxpool.Pool
}

func NewRemainRepo(db *pgxpool.Pool) *remainRepo {
	return &remainRepo{
		db: db,
	}
}

func (c *remainRepo) CreateRemain(req *models.CreateRemain) (string, error) {
	var (
		id         = uuid.NewString()
		totalPrice = req.Count * req.Price
		query      string
	)

	query = `
		INSERT INTO "remain"(
			"id", 
			"branch_id",
			"category_id",
			"name",
			"price",
			"barcode",
			"count",
			"total_price",
			"created_at" )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())`

	_, err := c.db.Exec(context.Background(), query,
		id,
		req.Branch_id,
		req.Category_id,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		totalPrice,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *remainRepo) GetRemain(req *models.RemainIdRequest) (*models.Remain, error) {
	query := `
		SELECT
		    "id",
		    "branch_id",
		    "category_id",
		    "name",
		    "price",
		    "barcode",
		    "count",
		    "total_price",
		    "created_at",
			   "updated_at"
		FROM "remain"
		WHERE id = $1
	`
	var (
		createdAt  time.Time
		updatedAt  sql.NullTime
		totalPrice float64
	)

	rem := models.Remain{}
	err := c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&rem.ID,
		&rem.Branch_id,
		&rem.Category_id,
		&rem.Name,
		&rem.Price,
		&rem.Barcode,
		&rem.Count,
		&totalPrice,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("Remain not found")
	}
	rem.TotalPrice = totalPrice
	rem.CreatedAt = createdAt.Format(time.RFC3339)
	if updatedAt.Valid {
		rem.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &rem, nil
}

func (c *remainRepo) GetAllRemain(req *models.GetAllRemainRequest) (*models.GetAllRemainResponse, error) {
	params := make(map[string]interface{})
	resp := &models.GetAllRemainResponse{}

	resp.Remainings = make([]models.Remain, 0)

	filter := " WHERE true "
	query := `
		SELECT
			COUNT(*) OVER(),
			"id",
			"branch_id",
			"category_id",
			"name",
			"price",
			"barcode",
			"count",
			"total_price",
			"created_at",
			"updated_at"
		FROM "remain"
	`
	if req.Search != "" {
		filter += ` AND ("category_id" ILIKE '%' || :search || '%' OR "branch_id" = :search || '%' OR "barcode" = :search) `
		params["search"] = req.Search
	}
	offset := (req.Page - 1) * req.Limit
	params["limit"] = req.Limit
	params["offset"] = offset

	query = query + filter + " ORDER BY created_at DESC OFFSET :offset LIMIT :limit "
	rquery, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := c.db.Query(context.Background(), rquery, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var (
		totalPrice float64
		createdAt  time.Time
		updatedAt  sql.NullTime
	)

	count := 0
	for rows.Next() {
		var rem models.Remain
		count++
		err := rows.Scan(
			&count,
			&rem.ID,
			&rem.Branch_id,
			&rem.Category_id,
			&rem.Name,
			&rem.Price,
			&rem.Barcode,
			&rem.Count,
			&totalPrice,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		rem.TotalPrice = totalPrice
		rem.CreatedAt = createdAt.Format(time.RFC3339)
		if updatedAt.Valid {
			rem.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
		}
		resp.Remainings = append(resp.Remainings, rem)
	}

	resp.Count = count
	return resp, nil
}
func (c *remainRepo) UpdateRemain(req *models.UpdateRemain) (string, error) {

	query := `UPDATE remain 
	            SET  branch_id = $1, 
				     category_id = $2,
					 name=$3,
					 price=$4,
					 barcode=$5,
					 count=$6,
					 total_price=$7, 
					 updated_at = NOW() 
					 WHERE id = $8 RETURNING id`

	result, err := c.db.Exec(context.Background(), query, req.Branch_id, req.Category_id, req.Name, req.Price, req.Barcode, req.Count, req.TotalPrice, req.ID)
	if err != nil {
		return "Error Update Remain", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Remain not found")
	}

	return req.ID, nil
}

func (c *remainRepo) DeleteRemain(req *models.RemainIdRequest) (resp string, err error) {
	query := `DELETE FROM remain 
	            WHERE id = $1 RETURNING id`

	result, err := c.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "Error from Delete Remain", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Remain not found")
	}

	return req.Id, nil
}
