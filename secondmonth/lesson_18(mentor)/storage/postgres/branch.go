package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/helper"
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
	id := uuid.NewString()
	yearNow := time.Now().Year()
	year := yearNow - req.FoundedAt

	query := `
		INSERT INTO branches(id, name, address, year, founded_at, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		year,
		req.FoundedAt,
	)

	if err != nil {
		return "Error CreateBranch", err
	}

	return id, nil
}

func (b *branchRepo) GetBranch(req *models.IdRequest) (resp *models.Branch, err error) {
	query := `SELECT id, name, address, year, founded_at, created_at, updated_at FROM branches WHERE id=$1`
	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	branch := models.Branch{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.Year,
		&branch.FoundedAt,
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

func (b *branchRepo) UpdateBranch(req *models.Branch) (string, error) {
	yearNow := time.Now().Year()
	year := yearNow - req.FoundedAt
	query := `UPDATE branches SET name = $1, address = $2, year = $3, founded_at = $4, updated_at = NOW() WHERE id = $5 RETURNING id`
	result, err := b.db.Exec(context.Background(), query, req.Name, req.Address, year, req.FoundedAt, req.ID)
	if err != nil {
		return "Error Update Branch", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch not found")
	}

	return req.ID, nil
}

func (b *branchRepo) GetAllBranch(req *models.GetAllBranchRequest) (*models.GetAllBranch, error) {
	params := make(map[string]interface{})
	filter := ""
	offset := (req.Page - 1) * req.Limit
	createdAt := time.Time{}
	updatedAt := sql.NullTime{}

	s := `SELECT id, name, address, year, founded_at, created_at, updated_at FROM branches order by created_at desc`

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
		return nil, err
	}
	defer rows.Close()

	resp := &models.GetAllBranch{}
	count := 0
	for rows.Next() {
		var branch models.Branch
		count++
		err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.Year,
			&branch.FoundedAt,
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

func (b *branchRepo) DeleteBranch(req *models.IdRequest) (resp string, err error) {
	query := `DELETE FROM branches WHERE id = $1 RETURNING id`

	result, err := b.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "Error from Delete Branch", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch not found")
	}

	return req.Id, nil
}
