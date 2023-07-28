package main

import "fmt"

type Person struct {
	FirstName string
	Lastname  string
	Job       string
	Age       int
	Id        int
}

var persons = []Person{}

func main() {
	// create a new person
	alex := Person{
		FirstName: "Alex",
		Lastname:  "Jack",
		Job:       "developer",
		Age:       76,
	}
	CreatePerson(alex)
	fmt.Println(persons)

	// read a person by id
	p, err := ReadPersonById(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p)
	}

	// read all persons
	ps := ReadAllPersons()
	fmt.Println(ps)

	// update a person by id
	updatedPerson := Person{
		FirstName: "John",
		Lastname:  "Doe",
		Job:       "engineer",
		Age:       35,
		Id:        1,
	}
	err = UpdatePersonById(updatedPerson)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(persons)
	}

	// delete a person by id
	err = DeletePersonById(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(persons)
	}
}

// CreatePerson adds a new person to the persons slice
func CreatePerson(newPerson Person) string {
	newPerson.Id = len(persons) + 1
	persons = append(persons, newPerson)
	return "successfully added"
}

// ReadPersonById returns the person with the specified id
func ReadPersonById(id int) (Person, error) {
	for _, p := range persons {
		if p.Id == id {
			return p, nil
		}
	}
	return Person{}, fmt.Errorf("person with id %d not found", id)
}

// ReadAllPersons returns the slice of all persons
func ReadAllPersons() []Person {
	return persons
}

// UpdatePersonById updates the person with the specified id
func UpdatePersonById(updatedPerson Person) error {
	for i, p := range persons {
		if p.Id == updatedPerson.Id {
			persons[i] = updatedPerson
			return nil
		}
	}
	return fmt.Errorf("person with id %d not found", updatedPerson.Id)
}

// DeletePersonById deletes the person with the specified id
func DeletePersonById(id int) error {
	for i, p := range persons {
		if p.Id == id {
			persons = append(persons[:i], persons[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("person with id %d not found", id)
}
