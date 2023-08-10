package memory

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"project/models"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
}

func NewBranchRepo(fn string) *branchRepo {
	return &branchRepo{fileName: fn}
}

func (b *branchRepo) CreateBranch(req models.CreateBranch) (string, error) {

	branches, err := b.read()
	if err != nil {
		return "", err
	}
	id := uuid.NewString()
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	branches = append(branches, models.Branch{
		Id:        id,
		Name:      req.Name,
		FoundedAt: req.FoundedAt,
		Address:   req.Address,
		CreatedAt: createdAt,
	})
	err = b.write(branches)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (b *branchRepo) UpdateBranch(req models.Branch) (msg string, err error) {
	branches, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range branches {
		if v.Id == req.Id {
			branches[i] = req
			msg = "Updated succesfully"
			err = b.write(branches)
			if err != nil {
				return "", nil
			}

		}
	}
	return "", errors.New("Not found")
}

func (b *branchRepo) GetBranch(req models.IdRequest) (resp models.Branch, err error) {
	branches, err := b.read()
	if err != nil {
		return models.Branch{}, err
	}

	for _, v := range branches {
		if v.Id == req.Id {
			foundedAt, _ := strconv.Atoi(v.FoundedAt[:4])
			v.Year = time.Now().Year() - foundedAt
			return v, nil
		}
	}
	return models.Branch{}, errors.New("Not found")
}

func (b *branchRepo) GetAllBranch(req models.GetAllBranchRequest) (resp models.GetAllBranch, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	branches, err := b.read()
	if err != nil {
		return models.GetAllBranch{}, err
	}

	if start >= len(branches) {
		resp.Branches = []models.Branch{}
		resp.Count = len(branches)
		return resp, nil
	} else if end > len(branches) {
		return models.GetAllBranch{
			Branches: branches[start:],
			Count:    len(branches),
		}, nil
	}

	return models.GetAllBranch{
		Branches: branches[start:end],
		Count:    len(branches),
	}, nil
}
func (b *branchRepo) DeleteBranch(req models.IdRequest) (msg string, err error) {

	branches, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range branches {
		if v.Id == req.Id {
			if len(branches) == 1 {
				branches = nil
			} else {
				branches = append(branches[:i], branches[i+1:]...)
			}
			msg = "Updated succesfully"
			err = b.write(branches)
			if err != nil {
				return "", nil
			}

		}
	}
	return "", errors.New("Not found")
}

func (b *branchRepo) read() ([]models.Branch, error) {
	var (
		branches []models.Branch
	)

	data, err := os.ReadFile(b.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)

		return nil, err
	}

	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return branches, nil
}

func (b *branchRepo) write(branches []models.Branch) error {

	body, err := json.Marshal(branches)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
