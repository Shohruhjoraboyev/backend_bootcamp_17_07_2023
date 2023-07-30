package main

import (
	"dars_6/cruds"
	"fmt"
)

func main() {
	branches := cruds.Branches{Data: make([]cruds.Branch, 0)}

	branches.Create(cruds.Branch{Name: "first", Address: "Yunusobod"})
	branches.Create(cruds.Branch{Name: "second", Address: "Olmazor"})
	branches.Create(cruds.Branch{Name: "third", Address: "Chilonzor"})

	for _, v := range branches.Data {
		fmt.Println(v)
	}
}
