package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{
		Nome:  "Vinicius",
		Idade: 22,
	}

	fmt.Println(pessoa)
}
