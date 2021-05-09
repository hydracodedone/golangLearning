package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type ProductService struct {
}

func (productSer *ProductService) GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{
		ProductList: []*ProductInfoList{
			&ProductInfoList{
				ProductId:    1,
				ProductPrice: 100,
				Time:         &timestamp.Timestamp{Seconds: time.Now().Unix()},
			},
			&ProductInfoList{
				ProductId:    2,
				ProductPrice: 200,
				Time:         &timestamp.Timestamp{Seconds: time.Now().Unix()},
			},
		},
	}, nil
}
