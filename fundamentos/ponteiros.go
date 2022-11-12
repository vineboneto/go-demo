package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) setNome(nome string) {
	p.Nome = nome
	fmt.Println(nome)
}

func main() {
	// nome := "Vinicius"
	// var nome2 *string

	// nome2 = &nome
	// *nome2 = "Teste"

	// fmt.Println(&nome == nome2)
	// fmt.Println("Conteúdo: ", nome)
	// fmt.Println("Apontamento memória: ", nome2)
	// fmt.Println("Conteúdo do apontamento: ", *nome2)

	pessoa := Pessoa{
		Nome:  "Vinicius",
		Idade: 22,
	}

	pessoa.setNome("José")
	fmt.Println(pessoa.Nome)

}
