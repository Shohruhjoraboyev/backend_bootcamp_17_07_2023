package main

import (
	"fmt"
	"time"
)

type Discount struct {
	Id              int    //unique id
	Name            string //
	Type            string // fixed  || percent
	DiscountType    string // discount || allowance
	Value           int
	DiscountObject  string   //order || product
	Sources         []string // website,bot,app,admin-panel
	PaymentTypes    []string // cash, payme, click,uzum
	ClientsOrderNum int      //1,2,....
	Priority        int
	Active          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Discounts struct {
	Data []Discount
}

func main() {
	discounts := Discounts{Data: make([]Discount, 0)}
	discounts.Create(Discount{
		Name:           "first discount",
		Type:           "percent",
		DiscountType:   "discount",
		Value:          15,
		DiscountObject: "order",
		Sources:        []string{"bot"},
		PaymentTypes:   []string{"cash", "click"},
		Priority:       1,
		Active:         true,
	})

	fmt.Println(discounts.Delete(0))

	res, err := discounts.GetByID(0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

func (d *Discounts) Create(newDiscount Discount) string {
	newDiscount.Id = len(d.Data)
	newDiscount.CreatedAt = time.Now()
	d.Data = append(d.Data, newDiscount)
	return "successfully added"
}
func (d *Discounts) Update(changedDiscount Discount) string {
	for i, v := range d.Data {
		if v.Id == changedDiscount.Id {
			d.Data[i].Name = changedDiscount.Name
			d.Data[i].Type = changedDiscount.Type
			d.Data[i].DiscountType = changedDiscount.DiscountType
			d.Data[i].Value = changedDiscount.Value
			d.Data[i].DiscountObject = changedDiscount.DiscountObject
			d.Data[i].Sources = changedDiscount.Sources
			d.Data[i].PaymentTypes = changedDiscount.PaymentTypes
			d.Data[i].ClientsOrderNum = changedDiscount.ClientsOrderNum
			d.Data[i].Priority = changedDiscount.Priority
			d.Data[i].Active = changedDiscount.Active
			d.Data[i].UpdatedAt = time.Now()
		}
	}
	// d.Data = append(d.Data, newDiscount)
	return "successfully added"
}

func (d *Discounts) GetByID(id int) (*Discount, error) {
	for i, v := range d.Data {
		if id == v.Id {
			return &d.Data[i], nil
		}
	}
	return nil, fmt.Errorf("not found by id=%d", id)
}

func (d *Discounts) Delete(id int) string {
	for i, v := range d.Data {
		if id == v.Id {
			d.Data[i] = d.Data[len(d.Data)-1]
			d.Data = d.Data[:len(d.Data)-1]
			return "successfully deleted"
		}
	}
	return "data not found"
}

func GetByID()
