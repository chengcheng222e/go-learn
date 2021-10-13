package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Position string `json:"position"`
	Name     Name   `json:"name"`
}

type Name struct {
	FirstName string `json:"firstname"`
	Surname   string `json:"surname"`
}

func main() {
	name := Name{FirstName: "Dikxya", Surname: "Lhyaho"}
	employee := Employee{Position: "Senior Developer", Name: name}

	byteArray, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
