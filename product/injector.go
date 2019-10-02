package product

import "golang-mysql/utils"

func ProvideProductRepository() ProductRepository {
	return NewProductRepository(utils.GetDatabaseConnection())
}

func ProvideProductService() ProductServices {
	return NewProductServices(ProvideProductRepository())
}

func ProvideProductHandler() ProductHandler {
	return NewProductHandler(ProvideProductService())
}
