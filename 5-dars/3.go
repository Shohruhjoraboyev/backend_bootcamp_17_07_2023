package main

// import "fmt"

// type StaffTariff struct {
// 	ID            int
// 	Name          string
// 	Type          string // fixed, percent
// 	AmountForCash float64
// 	AmountForCard float64
// }

// type StaffTariffs struct {
// 	Data []StaffTariff
// }

// func main() {
// 	staftariffs := StaffTariffs{}
// 	staftariffs.Create(StaffTariff{Name: "oddiy", Type: "fixed", AmountForCash: 12, AmountForCard: 10})

// 	for _, v := range staftariffs.Data {
// 		fmt.Println(v)
// 	}
// }

// func (s *StaffTariffs) Create(newTariff StaffTariff) string {
// 	newTariff.ID = len(s.Data) + 1
// 	s.Data = append(s.Data, newTariff)
// 	return "created successfully"
// }
