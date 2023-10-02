package postgres

// import (
// 	"context"
// 	"projectWithGinAndSwagger/models"

// 	"github.com/jackc/pgx/v4/pgxpool"
// )

// type topStaffRepo struct {
// 	db *pgxpool.Pool
// }

// func NewTopStaffRepo(db *pgxpool.Pool) *topStaffRepo {
// 	return &topStaffRepo{
// 		db: db,
// 	}
// }
// func (s *staffRepo) GetTopStaffs(req *models.TopStaffRequest) (resp *models.TopStaffResponse, err error) {
// 	top_staff := models.TopStaffResponse{}
// 	top_staff.TopStaff = make([]*models.TopStaff, 0)
// 	queryS := `
// 	SELECT
//     st.name,
//     b.name AS Branch_Name,
//     SUM(s.price) AS total_sum,
//     st.staff_type
// FROM sales AS s
// JOIN branches AS b ON b.id = s.branch_id
// JOIN staffs AS st ON st.id = s.shop_assistant_id
// WHERE s.status = 'success'
// GROUP BY s.cashier_id, st.name, b.name,st.staff_type
// Order by total_sum;
// `
// 	/*
// 		rows, err := s.db.Query(context.Background(), q, pArr...)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer rows.Close()

// 		resp = &models.GetAllStaff{}
// 		count := 0
// 		for rows.Next() {
// 			var staff models.Staff
// 			count++
// 			err := rows.Scan(
// 				&staff.Id,
// 				&staff.BranchId,
// 				&staff.TariffId,
// 				&staff.TypeId,
// 				&staff.Name,
// 				&staff.Balance,
// 				&staff.CreatedAt,
// 				&staff.UpdatedAt,
// 			)

// 			if err != nil {
// 				return nil, err
// 			}
// 			resp.Staffes = append(resp.Staffes, staff)

// 	*/

// 	rowsCashier, err := s.db.Query(context.Background())
// 	queryC := `
// 	SELECT
//     st.name,
//     b.name AS Branch_Name,
//     SUM(s.price) AS total_sum,
//     st.staff_type
// FROM sales AS s
// JOIN branches AS b ON b.id = s.branch_id
// JOIN staffs AS st ON st.id = s.cashier_id
// WHERE s.status = 'success'
// GROUP BY s.cashier_id, st.name, b.name,st.staff_type
// Order by total_sum;
// `

// 	return resp, nil

// }
