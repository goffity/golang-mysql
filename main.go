package main

import (
	"github.com/labstack/echo"
	"golang-mysql/product"
)

func main() {
	e := echo.New()
	p := product.NewProductRoute(product.ProvideProductHandler())
	p.Initial(e)

	e.Logger.Fatal(e.Start(":1323"))
}
