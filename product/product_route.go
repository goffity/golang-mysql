package product

import (
	"github.com/labstack/echo"
)

type ProductRoute interface {
	Initial(e *echo.Echo)
}

type productRoute struct {
	Handle ProductHandler
}

func NewProductRoute(handler ProductHandler) ProductRoute {
	return &productRoute{Handle: handler}
}

func (p productRoute) Initial(e *echo.Echo) {
	productGroup := e.Group("/products")
	productGroup.GET("/", p.Handle.GetProducts)
	productGroup.GET("/:id", p.Handle.GetProduct)
}
