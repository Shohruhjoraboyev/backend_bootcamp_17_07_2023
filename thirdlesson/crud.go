package main

import (
	"fmt"
)

func main() {
	// Yangi PersonStore yaratish
	ps := newPersonStore()

	// Yangi Person yaratish va saqlash
	p1 := Person{FirstName: "Bobur", LastName: "Boburov"}
	id1 := ps.CreatePerson(p1)

	// Person o'qish
	p2, err := ps.ReadPersonById(id1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read person:", p2)

	// Update person
	p2.FirstName = "Akmal"
	err = ps.UpdatePerson(p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Updated person:", p2)

	// Personni o'chirish
	err = ps.DeletePersonById(id1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Deleted person with id", id1)
}

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

type PersonStore struct {
	persons map[int]Person
	nextId  int
}

func newPersonStore() *PersonStore {
	return &PersonStore{
		persons: make(map[int]Person),
		nextId:  1,
	}
}

func (ps *PersonStore) CreatePerson(p Person) int {
	id := ps.nextId
	p.Id = id
	ps.persons[id] = p
	ps.nextId++
	return id
}

func (ps *PersonStore) ReadPersonById(id int) (Person, error) {
	p, ok := ps.persons[id]
	if !ok {
		return Person{}, fmt.Errorf("Person with id %d not found", id)
	}
	return p, nil
}

func (ps *PersonStore) UpdatePerson(p Person) error {
	_, ok := ps.persons[p.Id]
	if !ok {
		return fmt.Errorf("Person with id %d not found", p.Id)
	}
	ps.persons[p.Id] = p
	return nil
}

func (ps *PersonStore) DeletePersonById(id int) error {
	_, ok := ps.persons[id]
	if !ok {
		return fmt.Errorf("Person with id %d not found", id)
	}
	delete(ps.persons, id)
	return nil
}
