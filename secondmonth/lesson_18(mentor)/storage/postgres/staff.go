package postgres

import (
	"context"
	"fmt"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

// ChangeBalance implements storage.StaffesI.
func (*staffRepo) ChangeBalance(*models.ChangeBalance) (string, error) {
	panic("unimplemented")
}

// Exists implements storage.StaffesI.
func (*staffRepo) Exists(models.ExistsReq) bool {
	panic("unimplemented")
}

// GetByLogin implements storage.StaffesI.
func (*staffRepo) GetByLogin(models.LoginRequest) (*models.Staff, error) {
	panic("unimplemented")
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (s *staffRepo) CreateStaff(req *models.CreateStaff) (string, error) {
	id := uuid.NewString()
	query := `
	INSERT INTO staffs(id,branch_id,tariff_id,staff_type,name,balance,created_at)
	VALUES($1, $2, $3, $4, $5, $6,now())
	`
	_, err := s.db.Exec(context.Background(), query,
		id,
		req.BranchId,
		req.TariffId,
		req.StaffType,
		req.Name,
		req.Balance,
	)
	if err != nil {
		return "Error Create Staff", err
	}
	return id, nil
}

func (s *staffRepo) GetStaff(req *models.IdRequest) (resp *models.Staff, err error) {
	query := `SELECT * FROM staffs WHERE id=$1`
	staff := models.Staff{}
	err = s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staff.Id,
		&staff.BranchId,
		&staff.TariffId,
		&staff.TypeId,
		&staff.Name,
		&staff.Balance,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("Staff not found")
	}
	return &staff, nil

}
func (s *staffRepo) GetAllStaff(req *models.GetAllStaffRequest) (resp *models.GetAllStaff, err error) {
	params := make(map[string]interface{})
	filter := ""
	offset := (req.Page - 1) * req.Limit
	staff := `SELECT * FROM staffs`
	if req.Name != "" {
		filter += ` WHERE name ILIKE '%' || :search || '%' `
		params["search"] = req.Name
	}

	limit := fmt.Sprintf(" LIMIT %d", req.Limit)
	offsetQ := fmt.Sprintf(" OFFSET %d", offset)
	query := staff + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := s.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp = &models.GetAllStaff{}
	count := 0
	for rows.Next() {
		var staff models.Staff
		count++
		err := rows.Scan(
			&staff.Id,
			&staff.BranchId,
			&staff.TariffId,
			&staff.TypeId,
			&staff.Name,
			&staff.Balance,
			&staff.CreatedAt,
			&staff.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.Staffes = append(resp.Staffes, staff)

	}
	resp.Count = count
	return resp, nil
}
func (s *staffRepo) UpdateStaff(req *models.Staff) (string, error) {
	query := `UPDATE staffs SET branch_id= $1,tariff_id=$2,type_id=$3,name=$4,balance=$5,updated_at=now(),WHERE id=$6 RETURNING id`
	result, err := s.db.Exec(context.Background(), query, req.BranchId, req.TariffId, req.TypeId, req.Name, req.Balance, req.UpdatedAt, req.Id)
	if err != nil {
		return "Error Update staff", nil
	}
	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Staff not found")
	}
	return req.Id, nil
}

func (s *staffRepo) DeleteStaff(req *models.IdRequest) (resp string, err error) {
	query := `DELETE FROM staffs WHERE id=$1 RETURNING id`
	result, err := s.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Staff not found")
	}
	return req.Id, nil
}
