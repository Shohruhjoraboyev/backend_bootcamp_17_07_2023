package main

import (
	"fmt"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Job       string
}

var personList []Person

func createPerson(newPerson Person) {
	highestID := 0
	for _, person := range personList {
		if person.ID > highestID {
			highestID = person.ID
		}
	}

	newPerson.ID = highestID + 1
	personList = append(personList, newPerson)
}

func readPersonByID(personID int) (Person, error) {
	for _, person := range personList {
		if person.ID == personID {
			return person, nil
		}
	}

	return Person{}, fmt.Errorf("Person with ID %d not found", personID)
}

func readAllPerson() []string {
	persons := []string{}
	fmt.Println("Initial list of persons:")
	for _, person := range personList {
		s := fmt.Sprintf("ID: %d, Firstname: %s, LastName: %s, Age: %d, Job: %s\n", person.ID, person.FirstName, person.LastName, person.Age, person.Job)
		persons = append(persons, s)
	}

	return persons
}

func updatePerson(updatedPerson Person) error {
	for i, person := range personList {
		if person.ID == updatedPerson.ID {
			personList[i] = updatedPerson
			return nil
		}
	}

	return fmt.Errorf("Person with ID %d not found", updatedPerson.ID)
}

func deletePerson(personID int) error {
	for i, person := range personList {
		if person.ID == personID {
			personList[i] = personList[len(personList)-1]
			personList = personList[:len(personList)-1]
			return nil
		}
	}

	return fmt.Errorf("Person with ID %d not found", personID)
}

func main() {
	createPerson(Person{ID: 1, FirstName: "Omadbek", LastName: "Qosimov", Age: 19, Job: "freelancer"})
	createPerson(Person{ID: 2, FirstName: "Samandar", LastName: "Toshtemirov", Age: 20, Job: "student"})
	createPerson(Person{ID: 3, FirstName: "Sherzod", LastName: "Qosimov", Age: 21, Job: "3d designer"})
	createPerson(Person{ID: 4, FirstName: "Zafar", LastName: "Dilmurodov", Age: 19, Job: "developer"})

	for _, person := range readAllPerson() {
		fmt.Print(person)
	}

	person, _ := readPersonByID(3)

	fmt.Printf("\nRead by ID: 3\nID: %d, Firstname: %s, LastName: %s, Age: %d, Job: %s\n", person.ID, person.FirstName, person.LastName, person.Age, person.Job)

	updatedPerson := Person{ID: 1, FirstName: "Omadbek", LastName: "Qosimov", Age: 20, Job: "student"}
	err := updatePerson(updatedPerson)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nUpdated list of persons:")
	for _, person := range personList {
		fmt.Printf("ID: %d, FullName: %s, Age: %d, Job: %s\n", person.ID, person.LastName+" "+person.FirstName, person.Age, person.Job)
	}

	deletedPersonID := 2
	err = deletePerson(deletedPersonID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nFinal list of persons:")
	for _, person := range personList {
		fmt.Printf("ID: %d, FullName: %s, Age: %d, Job: %s\n", person.ID, person.LastName+" "+person.FirstName, person.Age, person.Job)
	}
}
