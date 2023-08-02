package main

import (
	"errors"
	"fmt"
	"time"
)

type Field struct {
	Id               int
	BranchId         int
	ShopAssistant_Id int
	Cashier_id       int
	Price            float64
	Payment_type     string
	Status           string
	Client_name      string
	CreatedAt        string
}

type Fields struct {
	Data []Field
}

func (f *Fields) Create(newField Field) {
	newField.Id = len(f.Data) + 1
	newField.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	f.Data = append(f.Data, newField)
	fmt.Println("Successfully added new Field")

}

func (f *Fields) Update(updatedField Field) {
	for i, field := range f.Data {
		if field.Id == updatedField.Id {
			updatedField.CreatedAt = field.CreatedAt
			f.Data[i] = updatedField
			fmt.Println("Successfully updated Field")
			return
		}
	}
	fmt.Println("Field not found")
}

func (f *Fields) GetById(id int) (Field, error) {
	for _, field := range f.Data {
		if field.Id == id {
			return field, nil
		}
	}
	return Field{}, errors.New("Field not found")
}

func (f *Fields) GetAll(branchId, shopAssistant_Id, cashier_id int, payment_type, status string, client_name *string, priceFrom float64, priceTo *float64) []Field {
	result := make([]Field, 0)

	for _, field := range f.Data {
		if (branchId == 0 || field.BranchId == branchId) &&
			(shopAssistant_Id == 0 || field.ShopAssistant_Id == shopAssistant_Id) &&
			(cashier_id == 0 || field.Cashier_id == cashier_id) &&
			(status == "" || field.Status == status) &&
			(payment_type == "" || field.Payment_type == payment_type) &&
			(client_name == nil || field.Client_name == "" || contains(field.Client_name, *client_name)) &&
			(priceFrom == 0 || field.Price >= priceFrom) &&
			(priceTo == nil || *priceTo == 0 || field.Price <= *priceTo) {
			result = append(result, field)
		}
	}
	return result
}

func contains(s, substr string) bool {
	return len(substr) == 0 || len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}

func (f *Fields) Delete(id int) {
	for i, field := range f.Data {
		if field.Id == id {
			f.Data = append(f.Data[:i], f.Data[i+1:]...)
			fmt.Println("Successfully Deleted Field")
			return
		}
	}
	fmt.Println("Field not found")
}

func main() {
	fields := Fields{Data: make([]Field, 0)}

	//Create
	for i := 0; i < 5; i++ {
		fields.Create(Field{
			Id:               1,
			BranchId:         1,
			ShopAssistant_Id: 2,
			Cashier_id:       2,
			Price:            120.0,
			Payment_type:     "Card",
			Status:           "Success",
			Client_name:      "Sarvarbek",
		})
	}

	//Update
	fields.Update(Field{
		Id:               1,
		BranchId:         2,
		ShopAssistant_Id: 1,
		Cashier_id:       1,
		Price:            130.00,
		Payment_type:     "Cash",
		Status:           "Cancel",
		Client_name:      "Omadbek",
	})

	//GetById
	field, err := fields.GetById(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Field found: ", field)
	}

	//GetAll

	allField := fields.GetAll(1, 0, 0, "", "Success", nil, 0.0, nil)
	fmt.Println("All fields with status 'Success': ", allField)

	//Delete
	fields.Delete(3)
}
