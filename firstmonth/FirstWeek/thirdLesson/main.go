// package main

// import "fmt"

// type Person struct {
// 	FirstName string
// 	Lastname  string
// 	Job       string
// 	Age       int
// 	Id        int
// }

// var persons = []Person{}

// func main() {

// 	alex := Person{
// 		FirstName: "Alex",
// 		Lastname:  "Jack",
// 		Job:       "developer",
// 		Age:       76,
// 	}
// 	CreatePerson(alex)
// 	fmt.Println(persons)

// 	Delete(2)
// }
// func CreatePerson(newPerson Person) string {
// 	newPerson.Id = len(persons) + 1
// 	persons = append(persons, newPerson)
// 	return "succesfully added"
// }

