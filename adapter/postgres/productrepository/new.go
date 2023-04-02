package productrepository

import (
	"github.com/LeandroMelloo/clean-architecture-com-golang/adapter/postgres"
	"github.com/LeandroMelloo/clean-architecture-com-golang/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
