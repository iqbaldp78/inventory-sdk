package service

import (
	"context"
	"errors"
	// "fmt"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/crud"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
)

func (s *service) GetOnhandQTY(ctx context.Context, input *pb.Empty) (*pb.ResOnhandQTY, error) {
	obj := crud.OnhandQTY{}
	data := obj.Search(s.db)

	if data.IsInitial() {
		return nil, errors.New("Data not found")
	}

	temp := make([]*pb.OnhandQTY, 0)
	for _, row := range data.Data.([]crud.OnhandQTY) {
		tempRow := &pb.OnhandQTY{
			Sku:      row.SKU,
			Itemname: row.ItemName,
			Stock:    int64(row.Stock),
		}
		temp = append(temp, tempRow)
	}

	return &pb.ResOnhandQTY{
		Total: data.Total,
		Data:  temp,
	}, nil
}
