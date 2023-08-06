package memory

import (
	"errors"
	"lesson_15/models"
)

type saleRepo struct {
	sales []models.Sales
}

func NewSaleRepo() *saleRepo {
	return &saleRepo{sales: make([]models.Sales, 0)}
}

func (c *saleRepo) CreateSale(req models.CreateSales) (int, error) {
	var id int
	if len(c.sales) == 0 {
		id = 1
	} else {
		id = c.sales[len(c.sales)-1].Id + 1
	}

	c.sales = append(c.sales, models.Sales{
		Id:               id,
		Name:             req.Name,
		Price:            req.Price,
		Payment_Type:     req.Payment_Type,
		Status:           req.Status,
		Branch_id:        req.Branch_id,
		Shop_asissent_id: req.Shop_asissent_id,
		Cashier_id:       req.Cashier_id,
		Created_at:       req.Created_at,
	})
	return id, nil
}

func (c *saleRepo) UpdateSale(req models.Sales) (string, error) {
	for i, v := range c.sales {
		if v.Id == req.Id {
			c.sales[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not sale found with ID")
}

func (c *saleRepo) GetSale(req models.IdRequest) (resp models.Sales, err error) {
	for _, v := range c.sales {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Sales{}, errors.New("not found")
}

func (c *saleRepo) GetAllSale(req models.GetAllSalesRequest) (resp models.GetAllSalesResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(c.sales) {
		resp.Sales = []models.Sales{}
		resp.Count = len(c.sales)
		return
	} else if end > len(c.sales) {
		return models.GetAllSalesResponse{
			Sales: c.sales[start:],
			Count: len(c.sales),
		}, nil
	}

	return models.GetAllSalesResponse{
		Sales: c.sales[start:end],
		Count: len(c.sales),
	}, nil
}

func (c *saleRepo) DeleteSale(req models.IdRequest) (resp string, err error) {
	for i, v := range c.sales {
		if v.Id == req.Id {
			if i == len(c.sales)-1 {
				c.sales = c.sales[:i]
			} else {
				c.sales = append(c.sales[:i], c.sales[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
