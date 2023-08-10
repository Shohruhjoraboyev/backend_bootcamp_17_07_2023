package memory

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"project/models"
	"time"

	"github.com/google/uuid"
)

type salesRepo struct {
	fileName string
}

func NewSalesRepo(fn string) *salesRepo {
	return &salesRepo{fileName: fn}
}

func (s *salesRepo) CreateSale(req models.CreateSales) (string, error) {

	sales, err := s.read()
	if err != nil {
		return "", err
	}
	id := uuid.NewString()
	sales = append(sales, models.Sales{
		Id:              id,
		ClientName:      req.ClientName,
		Price:           req.Price,
		PaymentType:     req.PaymentType,
		Status:          req.Status,
		BranchId:        req.BranchId,
		Shop_asissentId: req.Shop_asissentId,
		CashierId:       req.CashierId,
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
	})
	err = s.write(sales)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *salesRepo) UpdateSale(req models.Sales) (msg string, err error) {
	sales, err := s.read()
	if err != nil {
		return "", err
	}
	for i, v := range sales {
		if v.Id == req.Id {
			sales[i] = req
			msg = "Updated succesfully"
			err = s.write(sales)
			if err != nil {
				return "", nil
			}
		}
	}
	return "", errors.New("Not found")
}

func (s *salesRepo) GetSale(req models.IdRequest) (resp models.Sales, err error) {
	sales, err := s.read()
	if err != nil {
		return models.Sales{}, err
	}
	for _, v := range sales {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Sales{}, errors.New("not found")
}

func (s *salesRepo) GetAllSale(req models.GetAllSalesRequest) (resp models.GetAllSalesResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	sales, err := s.read()
	if err != nil {
		return models.GetAllSalesResponse{}, err
	}
	if start > len(sales) {
		resp.Sales = []models.Sales{}
		resp.Count = len(sales)
		return
	} else if end > len(sales) {
		return models.GetAllSalesResponse{
			Sales: sales[start:],
			Count: len(sales),
		}, nil
	}

	return models.GetAllSalesResponse{
		Sales: sales[start:end],
		Count: len(sales),
	}, nil
}

func (s *salesRepo) DeleteSale(req models.IdRequest) (string, error) {
	sales, err := s.read()
	if err != nil {
		return "", err
	}
	for i, v := range sales {
		if v.Id == req.Id {
			if len(sales) == 1 {
				sales = nil
			} else {
				sales = append(sales[:i], sales[i+1:]...)
			}
			return "Deleted successfully", nil
		}
	}
	return "", errors.New("Not found")
}

func (s *salesRepo) read() ([]models.Sales, error) {
	var (
		sales []models.Sales
	)

	data, err := os.ReadFile(s.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &sales)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return sales, nil
}

func (s *salesRepo) write(sales []models.Sales) error {

	body, err := json.Marshal(sales)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
