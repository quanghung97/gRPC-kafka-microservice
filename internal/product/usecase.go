package product

import (
	"context"

	"github.com/quanghung97/gRPC-kafka-microservice/internal/models"
)

// UseCase Product
type UseCase interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, productID string) (*models.Product, error)
	Search(ctx context.Context, search string, page, size int64) ([]*models.Product, error)
}
