package usecase

import (
	"context"

	"github.com/quanghung97/gRPC-kafka-microservice/internal/models"
	"github.com/quanghung97/gRPC-kafka-microservice/internal/product"
	"github.com/quanghung97/gRPC-kafka-microservice/pkg/logger"
)

// productUC
type productUC struct {
	productRepo product.MongoRepository
	log         logger.Logger
}

func NewProductUC(productRepo product.MongoRepository, log logger.Logger) *productUC {
	return &productUC{productRepo: productRepo, log: log}
}

func (p *productUC) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	panic("implement me")
}

func (p *productUC) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	panic("implement me")
}

func (p *productUC) GetByID(ctx context.Context, productID string) (*models.Product, error) {
	panic("implement me")
}

func (p *productUC) Search(ctx context.Context, search string, page, size int64) ([]*models.Product, error) {
	panic("implement me")
}
