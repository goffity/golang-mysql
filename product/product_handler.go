package product

import (
	"github.com/labstack/echo"
	"golang-mysql/model"
	"net/http"
	"strconv"
)

type ProductHandler interface {
	GetProducts(c echo.Context) error
	GetProduct(c echo.Context) error
}

type productHandler struct {
	ProductServices ProductServices
}

func NewProductHandler(services ProductServices) ProductHandler {
	return &productHandler{ProductServices: services}
}

func (p productHandler) GetProducts(c echo.Context) error {
	products, err := p.ProductServices.GetProductAll()
	if err.Code != 0 {
		c.Logger().Errorf(err.Message)
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, products)
}

func (p productHandler) GetProduct(c echo.Context) error {
	productIdString := c.Param("id")

	var product Product
	var response model.ResponseMessage

	if productIdString != "" {
		productId, err := strconv.Atoi(productIdString)
		if productId == 0 && err != nil {
			response = model.ResponseMessage{Code: http.StatusBadRequest, Message: err.Error(), DisplayMessage: ""}
		} else if err != nil {
			product, response = p.ProductServices.GetProduct(productId)
		}
	} else {
		response = model.ResponseMessage{Code: http.StatusBadRequest, Message: "Incomplete parameter", DisplayMessage: "Incomplete parameter"}
	}
	if response.Code != 0 {
		return c.JSON(response.Code, response)
	} else {
		return c.JSON(http.StatusOK, product)
	}
}
