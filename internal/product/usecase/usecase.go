package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/quanghung97/gRPC-kafka-microservice/internal/models"
	"github.com/quanghung97/gRPC-kafka-microservice/internal/product"
	"github.com/quanghung97/gRPC-kafka-microservice/pkg/logger"
)

// productUC
type productUC struct {
	productRepo product.MongoRepository
	log         logger.Logger
}

// NewProductUC productUC constructor
func NewProductUC(productRepo product.MongoRepository, log logger.Logger) *productUC {
	return &productUC{productRepo: productRepo, log: log}
}

// Create Create new product
func (p *productUC) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUC.Create")
	defer span.Finish()
	return p.productRepo.Create(ctx, product)
}

func (p *productUC) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUC.Update")
	defer span.Finish()
	return p.productRepo.Update(ctx, product)
}

func (p *productUC) GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUC.GetByID")
	defer span.Finish()
	return p.productRepo.GetByID(ctx, productID)
}

func (p *productUC) Search(ctx context.Context, search string, page, size int64) ([]*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUC.Search")
	defer span.Finish()
	return p.productRepo.Search(ctx, search, page, size)
}
