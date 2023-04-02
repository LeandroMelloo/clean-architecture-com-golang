package productusecase

import "github.com/LeandroMelloo/clean-architecture-com-golang/core/domain"

type usecase struct {
	repository domain.ProductRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.ProductRepository) domain.ProductUseCase {
	return &usecase{
		repository: repository,
	}
}
