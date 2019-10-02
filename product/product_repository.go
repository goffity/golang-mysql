package product

import (
	"database/sql"
	"golang-mysql/model"
)

type ProductRepository interface {
	FindAll() ([]Product, model.ResponseMessage)
	Find(id int) (Product, model.ResponseMessage)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p productRepository) FindAll() ([]Product, model.ResponseMessage) {
	panic("implement me")
}

func (p productRepository) Find(id int) (Product, model.ResponseMessage) {
	panic("implement me")
}
