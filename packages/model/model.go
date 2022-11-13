package model

import (
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Products []*Product
}

func (list *Products) Add(product *Product) {
	list.Products = append(list.Products, product)
}

func NewProduct(name string) *Product {

	return &Product{
		ID:   uuid.NewV4().String(),
		Name: string(name),
	}

}
