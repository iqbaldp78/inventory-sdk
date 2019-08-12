package service

import (
	"context"
	"errors"
	// "fmt"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/crud"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
)

func (s *service) GetInbounds(ctx context.Context, empty *pb.Empty) (*pb.ResInbounds, error) {
	obj := crud.Inbound{}

	data := obj.SearchInbounds(s.db)
	if data.IsInitial() {
		return nil, errors.New("Data not found")
	}
	temp := make([]*pb.Inbound, 0)
	for _, row := range data.Data.([]crud.Inbound) {

		tempRow := &pb.Inbound{
			PoDate:            row.PoDate,
			Sku:               row.SKU,
			PoLine:            int64(row.PoLine),
			QtyPo:             int64(row.QtyPO),
			QtyReceive:        int64(row.QtyReceive),
			PurchasePrice:     float32(row.PurchasePrice),
			TotalInboundPrice: float32(row.TotalInboundPrice),
			KwitansiNum:       row.KwitansiNum,
			Noted:             row.Noted,
			ItemName:          row.ItemName,
		}
		temp = append(temp, tempRow)
	}

	return &pb.ResInbounds{
		Total: data.Total,
		Data:  temp,
	}, nil
}
