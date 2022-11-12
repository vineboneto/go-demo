package main

import "fmt"

func main() {
	var nome string      // declarando
	nome = "Vinicius2"   // atribuindo
	nome1 := "Vinicius3" // declarando e atribuindo (realiza a inferência de tipo)
	fmt.Println(nome, nome1)
}
