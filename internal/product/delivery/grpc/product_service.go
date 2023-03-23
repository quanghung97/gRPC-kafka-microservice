package grpc

import (
	"github.com/quanghung97/gRPC-kafka-microservice/internal/product"
	"github.com/quanghung97/gRPC-kafka-microservice/pkg/logger"
)

// productGRPCService gRPC Service
type productGRPCService struct {
	log       logger.Logger
	productUC product.UseCase
}
