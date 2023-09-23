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

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) CreateCategory(req *models.CreateCategory) (string, error) {

	var (
		id    = uuid.NewString()
		query string
	)

	query = `
		INSERT INTO "category"(
			"id", 
			"name",
			"parent_id",
			"created_at" )
		VALUES ($1, $2, $3,  NOW())`

	_, err := c.db.Exec(context.Background(), query,
		id,
		req.Name,
		helper.NewNullString(req.Parent_id),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *categoryRepo) GetCategory(req *models.CategoryIdRequest) (resp *models.Category, err error) {

	query := `
		SELECT
		   "id", 
		    "name",
		    "parent_id",
		    "created_at" 
			"updated_at" 
		FROM "category"
		WHERE id = $1
	`
	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	category := models.Category{}
	err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&category.ID,
		&category.Name,
		&category.Parent_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("category not found")
	}
	category.CreatedAt = createdAt.Format(time.RFC3339)
	if updatedAt.Valid {
		category.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &category, nil
}

func (c *categoryRepo) GetAllCategory(req *models.GetAllCategoryRequest) (*models.GetAllCategoryResponse, error) {
	params := make(map[string]interface{})
	filter := ""
	offset := (req.Page - 1) * req.Limit
	createdAt := time.Time{}
	updatedAt := sql.NullTime{}

	s := `SELECT 
		   "id", 
		    "name",
		    "parent_id",
		    "created_at" 
			"updated_at" 
				FROM category order by created_at desc`

	if req.Name != "" {
		filter += ` WHERE name ILIKE '%' || :search || '%' `
		params["search"] = req.Name
	}

	limit := fmt.Sprintf(" LIMIT %d", req.Limit)
	offsetQ := fmt.Sprintf(" OFFSET %d", offset)
	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp := &models.GetAllCategoryResponse{}
	count := 0
	for rows.Next() {
		var category models.Category
		count++
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Parent_id,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		category.CreatedAt = createdAt.Format(time.RFC3339)
		if updatedAt.Valid {
			category.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
		}
		resp.Categories = append(resp.Categories, category)
	}

	resp.Count = count
	return resp, nil
}

func (c *categoryRepo) UpdateCategory(req *models.UpdateCategory) (string, error) {

	query := `UPDATE category 
	            SET  name = $1, 
				     parent_id = $2, 
					 updated_at = NOW() 
					 WHERE id = $3 RETURNING id`

	result, err := c.db.Exec(context.Background(), query, req.Name, req.Parent_id, req.Id)
	if err != nil {
		return "Error Update Category", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("category not found")
	}

	return req.Id, nil
}

func (c *categoryRepo) DeleteCategory(req *models.CategoryIdRequest) (resp string, err error) {
	query := `DELETE FROM category 
	            WHERE id = $1 RETURNING id`

	result, err := c.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "Error from Delete Category", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("Category not found")
	}

	return req.Id, nil
}
