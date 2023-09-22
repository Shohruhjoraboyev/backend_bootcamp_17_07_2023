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

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) CreateBranch(req *models.CreateBranch) (string, error) {

	var (
		id    = uuid.NewString()
		query string
	)

	query = `
		INSERT INTO "branches"(
			"id", 
			"name",
			"address",
			"phone",
			"created_at" )
		VALUES ($1, $2, $3, $4, NOW())`

	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		req.Phone,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *branchRepo) GetBranch(req *models.BranchIdRequest) (resp *models.Branch, err error) {

	query := `
		SELECT
			"id", 
			"name",
			"address",
			"phone",
			"created_at",
			"updated_at" 
		FROM "branches"
		WHERE id = $1
	`
	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	branch := models.Branch{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.Phone,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("branch not found")
	}
	branch.CreatedAt = createdAt.Format(time.RFC3339)
	if updatedAt.Valid {
		branch.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &branch, nil
}

func (b *branchRepo) GetAllBranch(req *models.GetAllBranchRequest) (*models.GetAllBranchResponse, error) {
	params := make(map[string]interface{})
	filter := ""
	offset := (req.Page - 1) * req.Limit
	createdAt := time.Time{}
	updatedAt := sql.NullTime{}

	s := `SELECT 
	            id, 
				name, 
				address, 
				phone, 
				created_at, 
				updated_at 
				FROM branches order by created_at desc`

	if req.Name != "" {
		filter += ` WHERE name ILIKE '%' || :search || '%' `
		params["search"] = req.Name
	}

	limit := fmt.Sprintf(" LIMIT %d", req.Limit)
	offsetQ := fmt.Sprintf(" OFFSET %d", offset)
	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp := &models.GetAllBranchResponse{}
	count := 0
	for rows.Next() {
		var branch models.Branch
		count++
		err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.Phone,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		branch.CreatedAt = createdAt.Format(time.RFC3339)
		if updatedAt.Valid {
			branch.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
		}
		resp.Branches = append(resp.Branches, branch)
	}

	resp.Count = count
	return resp, nil
}

func (b *branchRepo) UpdateBranch(req *models.UpdateBranch) (string, error) {

	query := `UPDATE branches 
	            SET  name = $1, 
				     address = $2, 
					 phone = $3, 
					 updated_at = NOW() 
					 WHERE id = $4 RETURNING id`

	result, err := b.db.Exec(context.Background(), query, req.Name, req.Address, req.Phone, req.Id)
	if err != nil {
		return "Error Update Branch", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch not found")
	}

	return req.Id, nil
}

func (b *branchRepo) DeleteBranch(req *models.BranchIdRequest) (resp string, err error) {
	query := `DELETE FROM branches 
	            WHERE id = $1 RETURNING id`

	result, err := b.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "Error from Delete Branch", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch not found")
	}

	return req.Id, nil
}
