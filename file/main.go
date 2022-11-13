package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

type List struct {
	Persons []*Person
}

func (person *Person) changeName(name string) {
	person.Name = name
}

func (list *List) add(person ...*Person) {
	list.Persons = append(list.Persons, person...)
}

func (list *List) toJson() []byte {
	out, err := json.Marshal(list.Persons)
	check(err)
	return out
}

func (person *Person) toJson() []byte {
	out, err := json.Marshal(person)
	check(err)
	return out
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getContentJSON() []Person {
	if _, err := os.Stat("./data.json"); err == nil {
		content, err := os.ReadFile("./data.json")

		check(err)

		if len(content) == 0 {
			return []Person{}
		}

		var payload []Person

		errContent := json.Unmarshal(content, &payload)

		check(errContent)

		return payload

	}
	f, err := os.Create("data.json")

	f.Close()

	check(err)

	return getContentJSON()
}

func main() {
	content := getContentJSON()

	list := List{}

	for i := 0; i < len(content); i++ {
		list.add(&content[i])
	}

	person := Person{Name: "Vinicius", Age: 15}

	list.add(&person)

	os.WriteFile("./data.json", list.toJson(), 0644)

	fmt.Println("done")
}
