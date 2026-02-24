package usecase

import (
	"errors"
	"rest-api-golang/model"
	"rest-api-golang/repository"
)

var (
	ErrProductNotFound = errors.New("Produto não encontrado")
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(id_product int) error {
	rowsAffected, err := pu.repository.DeleteProductById(id_product)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}
