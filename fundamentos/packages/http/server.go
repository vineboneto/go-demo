package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vineboneto/go-demo/model"
)

type WeberServer struct {
	Products *model.Products
}

func NewWebServer() *WeberServer {
	return &WeberServer{}
}

func (w WeberServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.POST("/products", w.createProduct)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WeberServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products.Products)
}

func (w WeberServer) createProduct(c echo.Context) error {
	product := model.NewProduct("")
	if err := c.Bind(product); err != nil {
		return err
	}
	fmt.Println(product)
	w.Products.Add(product)
	return c.JSON(http.StatusCreated, product)
}
