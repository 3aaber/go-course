package main

import (
	"errors"
)

//Person godoc
type Person struct {
	Name    string `json:"name"`
	Number  string `json:"number"`
	Address string `json:"address"`
}

var personsList = []Person{}

func getAllPersons() []Person {
	return personsList
}

func getPersonByName(name string) (*Person, error) {
	for _, val := range personsList {
		if val.Name == name {
			return &val, nil
		}
	}
	return nil, errors.New("Person not found")
}

func addNewPersonByData(name, number, address string) (*Person, error) {

	person := Person{Name: name, Number: number, Address: address}
	personsList = append(personsList, person)
	return &person, nil
}

func addNewPersonObject(p Person) {
	personsList = append(personsList, p)
}
