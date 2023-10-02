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

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (s *staffRepo) CreateStaff(ctx context.Context, req *models.CreateStaff) (string, error) {
	id := uuid.NewString()
	hashPassword, err := helper.GeneratePasswordHash(req.Password)
	req.Password = string(hashPassword)
	query := `
	INSERT INTO staffs(id,branch_id,tariff_id,staff_type,name,balance,login,password,phone,created_at)
	VALUES($1, $2, $3, $4, $5, $6,$7,$8,$9,now())	`

	_, err = s.db.Exec(ctx, query,
		id,
		req.BranchId,
		req.TariffId,
		req.StaffType,
		req.Name,
		req.Balance,
		req.Login,
		req.Password,
		req.Phone,
	)
	if err != nil {
		return "Error Create Staff", err
	}
	return id, nil
}

func (s *staffRepo) GetStaff(ctx context.Context, req *models.IdRequest) (resp *models.Staff, err error) {
	query := `SELECT * FROM staffs WHERE id=$1`
	staff := models.Staff{}
	err = s.db.QueryRow(ctx, query, req.Id).Scan(
		&staff.Id,
		&staff.BranchId,
		&staff.TariffId,
		&staff.TypeId,
		&staff.Name,
		&staff.Balance,
		&staff.Login,
		&staff.Password,
		&staff.Phone,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("Staff not found")
	}
	return &staff, nil

}
func (s *staffRepo) GetAllStaff(ctx context.Context, req *models.GetAllStaffRequest) (resp *models.GetAllStaff, err error) {
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
	rows, err := s.db.Query(ctx, q, pArr...)
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
			&staff.Login,
			&staff.Password,
			&staff.Phone,
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
func (s *staffRepo) UpdateStaff(ctx context.Context, req *models.Staff) (string, error) {
	query := `UPDATE staffs SET branch_id= $1,tariff_id=$2,type_id=$3,name=$4,balance=$5,login=$6,password=$7,phone=$8,updated_at=now(),WHERE id=$9 RETURNING id`
	result, err := s.db.Exec(ctx, query, req.BranchId, req.TariffId, req.TypeId, req.Name, req.Balance, req.Login, req.Password, req.Phone, req.UpdatedAt, req.Id)
	if err != nil {
		return "Error Update staff", nil
	}
	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Staff not found")
	}
	return req.Id, nil
}

func (s *staffRepo) ChangePassword(ctx context.Context, req *models.ReqNewPassword) (string, error) {
	query := `Update  staffs SET password=$1,WHERE id=$2  RETURNING Id`
	result, err := s.db.Exec(ctx, query, req.Password, req.Id)
	if err != nil {
		return "Error Change staff's Password", nil
	}
	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Staff not found fro changing Password")
	}
	return req.Id, nil
}

func (s *staffRepo) DeleteStaff(ctx context.Context, req *models.IdRequest) (resp string, err error) {
	query := `DELETE FROM staffs WHERE id=$1 RETURNING id`
	result, err := s.db.Exec(ctx, query, req.Id)

	if err != nil {
		return "", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Staff not found")
	}
	return req.Id, nil
}

func (b *staffRepo) UpdateBalance(ctx context.Context, req models.UpdateBalanceRequest) (res string, err error) {
	tr, err := b.db.Begin(ctx)
	defer func() {
		if err != nil {
			tr.Rollback(ctx)
		} else {
			tr.Commit(ctx)
		}
	}()

	cashierBQ := `
        UPDATE staffs
        SET balance-balance+$2
        WHERE id=$1 `

	if req.TransactionType == "withdraw" {
		req.Cashier.Amount = -req.Cashier.Amount
		req.ShopAssisstant.Amount = -req.ShopAssisstant.Amount
	}
	_, err = tr.Exec(ctx, cashierBQ, req.Cashier.StaffId, req.Cashier.Amount)
	if err != nil {
		return "", err
	}

	cashierTrQ := `
	INSERT INTO transactions(
     id,
	 sale_id,
	 amount,
	 type,
	 source_type,
	 text
	 )`

	_, err = tr.Exec(ctx, cashierTrQ,
		uuid.NewString(),
		req.Cashier.StaffId,
		req.SaleId,
		req.Cashier.Amount,
		req.TransactionType,
		req.SourceType,
		req.Text,
	)
	if err != nil {
		return "", err
	}

	if req.ShopAssisstant.StaffId != "" {
		ShAss_Balance_Q := `
        UPDATE staffs
        SET balance-balance+$2
        WHERE id=$1 `

		_, err = tr.Exec(ctx, ShAss_Balance_Q, req.ShopAssisstant.StaffId, req.ShopAssisstant.Amount)
		if err != nil {
			return "", err
		}

		ShAss_Tr_Q := `
	INSERT INTO transactions(
     id,
	 sale_id,
	 amount,
	 type,
	 source_type,
	 text
	 )`

		_, err = tr.Exec(ctx, ShAss_Tr_Q,
			uuid.NewString(),
			req.ShopAssisstant.StaffId,
			req.SaleId,
			req.ShopAssisstant.Amount,
			req.TransactionType,
			req.SourceType,
			req.Text,
		)
		if err != nil {
			return "", err
		}
	}
	return "Updated Balance", err
}

func (b *staffRepo) GetByLogin(ctx context.Context, req *models.LoginRequest) (*models.Staff, error) {

	s := `SELECT
         s.id,
         s.name,
         b.name,
         t.name,
         s.typeid,
         s.balance
      FROM
         staffs AS s
         JOIN branches AS b ON b.id = s.branch_id
         JOIN staff_tarif AS t ON t.id = s.tariff_id
      WHERE
         s.login = $1`

	staff := models.Staff{}

	err := b.db.QueryRow(ctx, s, req.Login).Scan(
		&staff.Login,
		&staff.Password,
	)

	if err != nil {
		return &staff, err
	}
	return &staff, err

}
