package product

import (
	"golang-mysql/model"
)

type ProductServices interface {
	GetProductAll() ([]Product, model.ResponseMessage)
	GetProduct(id int) (Product, model.ResponseMessage)
}

type productServices struct {
	productRepository ProductRepository
}

func NewProductServices(productRepository ProductRepository) ProductServices {
	return &productServices{productRepository: productRepository}
}

func (p productServices) GetProductAll() ([]Product, model.ResponseMessage) {
	return p.productRepository.FindAll()
}

func (p productServices) GetProduct(id int) (Product, model.ResponseMessage) {
	return p.productRepository.Find(id)
}
