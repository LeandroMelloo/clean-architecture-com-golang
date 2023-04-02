package productusecase

import (
	"github.com/LeandroMelloo/clean-architecture-com-golang/core/domain"
	"github.com/LeandroMelloo/clean-architecture-com-golang/core/dto"
)

func (usecase usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := usecase.repository.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
