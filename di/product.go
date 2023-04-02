package di

import (
	"github.com/LeandroMelloo/clean-architecture-com-golang/adapter/http/productservice"
	"github.com/LeandroMelloo/clean-architecture-com-golang/adapter/postgres"
	"github.com/LeandroMelloo/clean-architecture-com-golang/adapter/postgres/productrepository"
	"github.com/LeandroMelloo/clean-architecture-com-golang/core/domain"
	"github.com/LeandroMelloo/clean-architecture-com-golang/core/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
