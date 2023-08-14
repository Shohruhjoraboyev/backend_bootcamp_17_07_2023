package memory

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"project/models"

	"github.com/google/uuid"
)

type staffRepo struct {
	fileName string
}

func NewStaffRepo(fn string) *staffRepo {
	return &staffRepo{fileName: fn}
}

func (s *staffRepo) CreateStaff(req models.CreateStaffRequest) (string, error) {
	id := uuid.NewString()
	staffs, err := s.read()
	if err != nil {
		return "", err
	}
	staffs = append(staffs, models.Staff{
		Id:       id,
		BranchId: req.BranchId,
		TariffId: req.TariffId,
		Type:     string(req.Type),
		Name:     req.Name,
		Balance:  req.Balance,
	})
	err = s.write(staffs)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *staffRepo) UpdateStaff(req models.Staff) (string, error) {
	staffs, err := s.read()
	if err != nil {
		return "", err
	}
	for i, v := range staffs {
		if v.Id == req.Id {
			staffs[i] = req
			return "Updated succesfully", nil
		}
	}
	err = s.write(staffs)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (s *staffRepo) GetStaff(req models.GetStaffByIdRequest) (resp models.Staff, err error) {
	staffs, err := s.read()
	if err != nil {
		return models.Staff{}, err
	}
	for _, v := range staffs {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Staff{}, errors.New("Not found")
}

func (s *staffRepo) GetAllStaff(req models.GetAllStaffRequest) (resp models.GetAllStaffResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	staffs, err := s.read()
	if err != nil {
		return models.GetAllStaffResponse{}, err
	}

	if start > len(staffs) {
		resp.Staffs = []models.Staff{}
		resp.Count = len(staffs)
		return
	} else if end > len(staffs) {
		return models.GetAllStaffResponse{
			Staffs: staffs[start:],
			Count:  len(staffs),
		}, nil
	}
	return models.GetAllStaffResponse{
		Staffs: staffs[start:end],
		Count:  len(staffs),
	}, nil

}

func (s *staffRepo) DeleteStaff(req models.GetStaffByIdRequest) (string, error) {
	staffs, err := s.read()
	if err != nil {
		return "", err
	}
	for i, v := range staffs {
		if v.Id == req.Id {
			if len(staffs) == 1 {
				staffs = nil
			} else {
				staffs = append(staffs[:i], staffs[i+1:]...)
			}
			return "Deleted succesfully", nil
		}
	}
	return "", errors.New("Not found")
}

func (s *staffRepo) read() ([]models.Staff, error) {
	var (
		staffs []models.Staff
	)

	data, err := os.ReadFile(s.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &staffs)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err

	}

	return staffs, nil
}

func (s *staffRepo) write(staffs []models.Staff) error {

	body, err := json.Marshal(staffs)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
