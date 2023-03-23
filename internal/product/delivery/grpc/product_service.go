package grpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/quanghung97/gRPC-kafka-microservice/internal/product"
	"github.com/quanghung97/gRPC-kafka-microservice/pkg/logger"
	productsService "github.com/quanghung97/gRPC-kafka-microservice/proto/product"
)

// productService gRPC Service
type productService struct {
	log       logger.Logger
	productUC product.UseCase
	validate  *validator.Validate
}

// NewProductService productService constructor
func NewProductService(log logger.Logger, productUC product.UseCase, validate *validator.Validate) *productService {
	return &productService{log: log, productUC: productUC, validate: validate}
}

func (p *productService) Create(ctx context.Context, req *productsService.UpdateReq) (*productsService.UpdateRes, error) {
	panic("implement me")
}

func (p *productService) Update(ctx context.Context, req *productsService.UpdateReq) (*productsService.UpdateRes, error) {
	panic("implement me")
}

func (p *productService) GetByID(ctx context.Context, req *productsService.GetByIDReq) (*productsService.GetByIDRes, error) {
	panic("implement me")
}

func (p *productService) Search(ctx context.Context, req *productsService.SearchReq) (*productsService.SearchRes, error) {
	panic("implement me")
}
