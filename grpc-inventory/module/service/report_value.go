package service

import (
	"context"
	"errors"
	// "fmt"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/crud"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
)

func (s *service) GetReportValue(ctx context.Context, empty *pb.Empty) (*pb.ResReportValue, error) {
	obj := crud.ReportValueHeader{}
	data := obj.SearchReportValue(s.db)
	if data.IsInitial() {
		return nil, errors.New("Data not found")
	}

	temp := make([]*pb.ReportValueHeader, 0)
	for _, row := range data.Data.([]crud.ReportValueHeader) {
		var sliceLines []*pb.ReportValueLines
		for _, val := range row.ReportValueLines {
			tempLines := &pb.ReportValueLines{
				AvgPruchasePrice:   val.AvgPruchasePrice,
				Itemname:           val.Itemname,
				QtyTotalInv:        val.QtyTotalInv,
				Sku:                val.Sku,
				TotalPurchasePrice: val.TotalPurchasePrice,
			}
			sliceLines = append(sliceLines, tempLines)
		}
		tempRow := &pb.ReportValueHeader{
			CountSku:         row.CountSku,
			SumStock:         row.SumStock,
			TotalValue:       row.TotalValue,
			PrintDate:        row.PrintDate,
			ReportValueLines: sliceLines,
		}
		temp = append(temp, tempRow)
	}
	return &pb.ResReportValue{
		Total: data.Total,
		Data:  temp,
	}, nil
}
