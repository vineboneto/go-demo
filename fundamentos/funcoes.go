package main

import (
	"errors"
	"fmt"
	"strings"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) andou() (string, error) {
	if strings.ToLower(p.Nome) != "vinicius" {
		return "", errors.New("Nome tem que ser vinicius!")
	}

	s := []string{"pessoa", strings.ToLower(p.Nome), "andou"}

	return strings.Join(s, " "), nil
}

func main() {
	pessoa := Pessoa{
		Nome:  "Vinicius",
		Idade: 22,
	}

	res, err := pessoa.andou()
	if err != nil {
		// Tratar erro
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
