package service

import (
	"context"
)

type ProductService struct {
}

func (productSer *ProductService) GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{ProductStock: 20}, nil
}
