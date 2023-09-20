package postgres

// import (
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
