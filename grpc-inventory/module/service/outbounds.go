package service

import (
	"context"
	"errors"
	// "fmt"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/crud"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
)

func (s *service) GetOutbounds(ctx context.Context, empty *pb.Empty) (*pb.ResOutbounds, error) {
	obj := crud.Outbound{}
	data := obj.SearchOutbounds(s.db)
	if data.IsInitial() {
		return nil, errors.New("Data not found")
	}

	temp := make([]*pb.Outbound, 0)
	for _, row := range data.Data.([]crud.Outbound) {
		tempRow := &pb.Outbound{
			OrderDate:          row.OrderDate,
			Sku:                row.Sku,
			Itemname:           row.Itemname,
			QtyOut:             row.QtyOut.Int64,
			SellingPrice:       row.SellingPrice,
			TotalOutboundPrice: row.TotalOutboundPrice,
			Noted:              row.Noted,
		}
		temp = append(temp, tempRow)
	}
	return &pb.ResOutbounds{
		Total: data.Total,
		Data:  temp,
	}, nil
}
