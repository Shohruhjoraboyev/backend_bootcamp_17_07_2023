package main

import "fmt"

func main() {

	Omadbek := Person{
		FirstName: "Omadbek",
		LastName:  "Jaloldinov",
		Job:       "Backend Developer",
		Age:       21,
	}

	for i := 0; i < 2; i++ {
		createPerson(Omadbek)
	}

	printer(containerArr, containerMap)
}

var containerMap = map[int]Person{}
var containerArr = []Person{}

type Person struct {
	FirstName string
	LastName  string
	Job       string
	Age       int
	Id        int
}

func createPerson(p Person) string {

	id := len(containerArr) + 1

	// array pusher
	containerArr = append(containerArr, Person{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Job:       p.Job,
		Age:       p.Age,
		Id:        id,
	})

	// map pusher
	containerMap[id] = Person{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Job:       p.Job,
		Age:       p.Age,
		Id:        id,
	}

	return "s"
}

func printer(arr []Person, mapp map[int]Person) {
	fmt.Println("arr", arr)
	fmt.Println("map", mapp)

}
