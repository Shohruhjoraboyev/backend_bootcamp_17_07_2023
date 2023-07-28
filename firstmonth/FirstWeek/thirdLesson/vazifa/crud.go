package main

import (
	"errors"
	"fmt"
)

type Person struct {
	FirstName string
	Lastname  string
	Job       string
	Age       int
	Id        int
}
type Persons struct {
	Pdata []Person
}

func main() {
	//Create
	persons := Persons{Pdata: make([]Person, 0)}
	for i := 0; i < 3; i++ {
		persons.Create(Person{
			FirstName: "Alex",
			Lastname:  "Jack",
			Job:       "developer",
			Age:       76,
		})
	}

	// Update
	fmt.Println(persons.Update(Person{FirstName: "Joanna"}))

	//GetbyId
	resp, err := persons.getById(3)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(resp)
	}

	//Delete
	fmt.Println(persons.Delete(2))

	//ReadALl
	persons.ReadAll()
}

// functions create ,getbyid,readall,update,delete

// Function of Create
func (p *Persons) Create(NewPerson Person) string {
	NewPerson.Id = len(p.Pdata) + 1
	p.Pdata = append(p.Pdata, NewPerson)
	return "Successfully added new Person"
}

// Update function
func (p *Persons) Update(changePerson Person) string {
	for i, v := range p.Pdata {
		if v.Id == changePerson.Id {
			p.Pdata[i].FirstName = changePerson.FirstName
			p.Pdata[i].Lastname = changePerson.Lastname
			p.Pdata[i].Job = changePerson.Job
			p.Pdata[i].Age = changePerson.Age
		}

	}
	return "Succesfully changed "
}

// getById
func (p *Persons) getById(id int) (Person, error) {
	for _, v := range p.Pdata {
		if v.Id == id {
			return v, nil
		}
	}
	return Person{}, errors.New("Not Found")
}

// delete function

func (p *Persons) Delete(id int) string {
	for i, v := range p.Pdata {
		if v.Id == id {
			p.Pdata = append(p.Pdata[:i], p.Pdata[i+1:]...)
			return "Deleted Succesfully"
		}
	}
	return "Not found this id"
}

// ReadAll function
func (p *Persons) ReadAll() {
	for _, person := range p.Pdata {
		fmt.Printf("ID: %d, Firstname: %s, LastName: %s, Age: %d, Job: %s\n", person.Id, person.FirstName, person.Lastname, person.Age, person.Job)
	}
}
