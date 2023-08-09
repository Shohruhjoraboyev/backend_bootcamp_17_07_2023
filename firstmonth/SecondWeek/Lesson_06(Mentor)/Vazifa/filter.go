package main

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	Lastname  string
	Job       string
	Age       int
	Id        int
}

var persons = map[int]Person{}

func main() {

	alex := Person{
		FirstName: "Azizbek",
		Lastname:  "Jack",
		Job:       "developer",
		Age:       76}

	for i := 0; i < 5; i++ {
		CreatePerson(alex)
	}

	UpdatedPerson := Person{
		FirstName: "Rustam",
		Lastname:  " ALiyev",
		Job:       "developer",
		Age:       26,
	}
	// fmt.Println(persons)

	Update(3, UpdatedPerson)

	// fmt.Println("After deleted: ")

	Delete(2)
	ReadbyId(3)
	// fmt.Println(persons)

	fmt.Println(ReadAll("", 12))

}
func CreatePerson(newPerson Person) string {
	newPerson.Id = len(persons) + 1

	persons[newPerson.Id] = newPerson
	return "Created new Person"
}

func Update(id int, updatePerson Person) string {

	_, ok := persons[id]
	if ok {
		updatePerson.Id = id
		persons[id] = updatePerson
		return "Updated"
	} else {
		return "Not found this id's person"
	}

}
func Delete(id int) string {

	_, ok := persons[id]
	if ok {
		delete(persons, id)
		return "Deleted successfully"
	} else {
		return "not found this id"

	}
}

func ReadbyId(id int) (Person, string) {
	value, ok := persons[id]
	if ok {
		return value, "Found"
	} else {
		return value, "Not Found"
	}

}
func ReadAll(nameFilter string, ageFilter int) ([]Person, string) {
	var result []Person
	for _, person := range persons {
		if (nameFilter == "" || strings.Contains(person.FirstName, nameFilter)) && (ageFilter == 0 || person.Age == ageFilter) {
			result = append(result, person)
		}
	}
	if len(result) != 0 {
		return result, "All data"
	} else {
		return result, "Not found"
	}
}
