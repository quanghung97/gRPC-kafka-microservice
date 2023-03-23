package grpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"github.com/quanghung97/gRPC-kafka-microservice/internal/models"
	"github.com/quanghung97/gRPC-kafka-microservice/internal/product"
	grpcErrors "github.com/quanghung97/gRPC-kafka-microservice/pkg/grpc_errors"
	"github.com/quanghung97/gRPC-kafka-microservice/pkg/logger"
	productsService "github.com/quanghung97/gRPC-kafka-microservice/proto/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// productService gRPC Service
type productService struct {
	log       logger.Logger
	productUC product.UseCase
	validate  *validator.Validate
	productsService.ProductsServiceServer
}

// NewProductService productService constructor
func NewProductService(log logger.Logger, productUC product.UseCase, validate *validator.Validate) *productService {
	return &productService{log: log, productUC: productUC, validate: validate}
}

func (p *productService) Create(ctx context.Context, req *productsService.CreateReq) (*productsService.CreateRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Create")
	defer span.Finish()

	catID, err := primitive.ObjectIDFromHex(req.GetCategoryID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	prod := &models.Product{
		CategoryID:  catID,
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		ImageURL:    &req.ImageURL,
		Photos:      req.GetPhotos(),
		Quantity:    req.GetQuantity(),
		Rating:      int(req.GetRating()),
	}

	if err := p.validate.StructCtx(ctx, prod); err != nil {
		p.log.Errorf("validate.StructCtx: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	created, err := p.productUC.Create(ctx, prod)
	if err != nil {
		p.log.Errorf("productUC.Create: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	return &productsService.CreateRes{Product: created.ToProto()}, nil
}

func (p *productService) Update(ctx context.Context, req *productsService.UpdateReq) (*productsService.UpdateRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Update")
	defer span.Finish()

	prodID, err := primitive.ObjectIDFromHex(req.GetProductID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}
	catID, err := primitive.ObjectIDFromHex(req.GetCategoryID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	prod := &models.Product{
		ProductID:   prodID,
		CategoryID:  catID,
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		ImageURL:    &req.ImageURL,
		Photos:      req.GetPhotos(),
		Quantity:    req.GetQuantity(),
		Rating:      int(req.GetRating()),
	}

	if err := p.validate.StructCtx(ctx, prod); err != nil {
		p.log.Errorf("validate.StructCtx: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	update, err := p.productUC.Update(ctx, prod)
	if err != nil {
		p.log.Errorf("productUC.Update: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	return &productsService.UpdateRes{Product: update.ToProto()}, nil
}

func (p *productService) GetByID(ctx context.Context, req *productsService.GetByIDReq) (*productsService.GetByIDRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.GetByID")
	defer span.Finish()

	prodID, err := primitive.ObjectIDFromHex(req.GetProductID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	prod, err := p.productUC.GetByID(ctx, prodID)
	if err != nil {
		p.log.Errorf("productUC.GetByID: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	return &productsService.GetByIDRes{Product: prod.ToProto()}, nil
}

func (p *productService) Search(ctx context.Context, req *productsService.SearchReq) (*productsService.SearchRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Search")
	defer span.Finish()

	products, err := p.productUC.Search(ctx, req.GetSearch(), req.GetPage(), req.GetSize())
	if err != nil {
		p.log.Errorf("productUC.Search: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	p.log.Infof("PRODUCTS: %-v", products)

	return nil, nil
}
