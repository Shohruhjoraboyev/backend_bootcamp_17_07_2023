package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/helper"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffTarifRepo struct {
	db *pgxpool.Pool
}

func NewStaffTarifRepo(db *pgxpool.Pool) *staffTarifRepo {
	return &staffTarifRepo{db: db}
}

func (s *staffTarifRepo) CreateStaffTarif(ctx context.Context, req *models.CreateStaffTarif) (string, error) {
	id := uuid.NewString()

	query := `
    INSERT INTO "staff_tarif" 
    ("id", "name", "type", "amount_for_cash", "amount_for_card", "created_at")
    VALUES 
    ($1, $2, $3, $4, $5, NOW())
  `

	_, err := s.db.Exec(ctx, query,
		id,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create staff tariff: %w", err)
	}

	return id, nil
}

func (s *staffTarifRepo) GetStaffTarif(ctx context.Context, req *models.IdRequest) (resp *models.StaffTarif, err error) {
	query := `
    SELECT  "id", "name", "type", "amount_for_cash", "amount_for_card", "created_at", "updated_at"
    FROM "staff_tarif" WHERE "id" = $1
  `
	var (
		created_at time.Time
		updated_at sql.NullTime
	)

	tariff := models.StaffTarif{}
	err = s.db.QueryRow(ctx, query, req.Id).Scan(
		&tariff.Id,
		&tariff.Name,
		&tariff.Type,
		&tariff.AmountForCard,
		&tariff.AmountForCash,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("staff tariff not found")
		}
		return nil, fmt.Errorf("failed to get staff tariff: %w", err)
	}

	tariff.CreatedAt = created_at.Format(time.RFC3339)
	if updated_at.Valid {
		tariff.UpdatedAt = updated_at.Time.Format(time.RFC3339)
	}

	return &tariff, nil
}
func (s *staffTarifRepo) GetAllStaffTarif(ctx context.Context, req *models.GetAllStaffTarifRequest) (resp *models.GetAllStaffTarif, err error) {
	params := make(map[string]interface{})
	filter := ""
	created_at := time.Time{}
	updated_at := sql.NullTime{}

	sekect := `
    SELECT
    "id", "name", "type", "amount_for_cash", "amount_for_card", "created_at", "updated_at"
    FROM "staff_tarif"
  `

	if req.Name != "" {
		filter += ` WHERE name ILIKE '%' || :search || '%' `
		params["search"] = req.Name
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

	query := sekect + filter + " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := s.db.Query(ctx, q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp = &models.GetAllStaffTarif{}
	count := 0
	for rows.Next() {
		var tariff models.StaffTarif
		count++
		err := rows.Scan(
			&tariff.Id,
			&tariff.Name,
			&tariff.Type,
			&tariff.AmountForCard,
			&tariff.AmountForCash,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tariff.CreatedAt = created_at.Format(time.RFC3339)
		if updated_at.Valid {
			tariff.UpdatedAt = updated_at.Time.Format(time.RFC3339)
		}
		resp.StaffTarifs = append(resp.StaffTarifs, tariff)
	}

	resp.Count = count
	return resp, nil
}

func (s *staffTarifRepo) UpdateStaffTarif(ctx context.Context, req *models.StaffTarif) (string, error) {
	query := `
    UPDATE "staff_tarif"
    SET "name" = $1, "type" = $2, "amount_for_cash" = $3, "amount_for_card" = $4, "updated_at" = NOW()
    WHERE "id" = $5
  `

	result, err := s.db.Exec(ctx, query,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
		req.Id,
	)
	if err != nil {
		return "", fmt.Errorf("failed to update staff tariff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff tariff with ID %s not found", req.Id)
	}

	return req.Id, nil
}

func (s *staffTarifRepo) DeleteStaffTarif(ctx context.Context, req *models.IdRequest) (string, error) {
	query := `
	  DELETE FROM staff_tarif
	  WHERE "id" = $1
	`

	result, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete staff tariff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff tariff with ID %s not found", req.Id)
	}

	return req.Id, nil
}
