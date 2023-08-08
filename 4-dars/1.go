package main

// import (
// 	"fmt"
// 	"strings"
// )

// type Person struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	Job       string
// }

// var personList []Person

// func createPerson(newPerson Person) {
// 	highestID := 0
// 	for _, person := range personList {
// 		if person.ID > highestID {
// 			highestID = person.ID
// 		}
// 	}

// 	newPerson.ID = highestID + 1
// 	personList = append(personList, newPerson)
// }

// func readAllPerson(age int, name string) []string {
// 	persons := []string{}
// 	fmt.Println("Initial list of persons:")
// 	for _, person := range personList {
// 		if person.Age == age && strings.Contains(person.FirstName, name) {
// 			s := fmt.Sprintf("ID: %d, Firstname: %s, LastName: %s, Age: %d, Job: %s\n", person.ID, person.FirstName, person.LastName, person.Age, person.Job)
// 			persons = append(persons, s)
// 		}
// 	}

// 	return persons
// }

// func main() {
// 	createPerson(Person{ID: 1, FirstName: "Omadbek", LastName: "Qosimov", Age: 19, Job: "freelancer"})
// 	createPerson(Person{ID: 2, FirstName: "Samandar", LastName: "Toshtemirov", Age: 20, Job: "student"})
// 	createPerson(Person{ID: 3, FirstName: "Sherzod", LastName: "Qosimov", Age: 21, Job: "3d designer"})
// 	createPerson(Person{ID: 4, FirstName: "Zafar", LastName: "Dilmurodov", Age: 19, Job: "developer"})

// 	fmt.Println("Filter: age=19,name:bek")
// 	for _, person := range readAllPerson(19, "a") {
// 		fmt.Print(person)
// 	}
// }
