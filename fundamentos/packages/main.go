package main

import (
	"github.com/vineboneto/go-demo/http"
	"github.com/vineboneto/go-demo/model"
)

func main() {
	produto1 := model.NewProduct("Carrinho")
	produto2 := model.NewProduct("Boneca")

	products := model.Products{}

	products.Add(produto1)
	products.Add(produto2)

	server := http.NewWebServer()
	server.Products = &products

	server.Serve()
}
