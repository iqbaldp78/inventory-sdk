package service

import (
	"context"
	"errors"
	// "fmt"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/crud"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
)

func (s *service) GetReportSales(ctx context.Context, param *pb.ParamDate) (*pb.ResReportSalesHeader, error) {
	obj := crud.ReportSalesHeader{}
	data := obj.SearchReportSales(s.db, param.DateFrom, param.DateTo)
	if data.IsInitial() {
		return nil, errors.New("Data not found")
	}

	temp := make([]*pb.ReportSalesHeader, 0)
	for _, row := range data.Data.([]crud.ReportSalesHeader) {
		var sliceLines []*pb.ReportSalesLines
		for _, val := range row.ReportSalesLines {
			tempLines := &pb.ReportSalesLines{
				OrderNum:      val.OrderNum,
				OrderDate:     val.OrderDate,
				Sku:           val.Sku,
				ItemName:      val.Itemname,
				QtySo:         val.QtySo,
				SellingPrice:  val.SellingPrice,
				TotalPrice:    val.TotalPrice,
				PurchasePrice: val.PurchasePrice,
				Laba:          val.Laba,
			}
			sliceLines = append(sliceLines, tempLines)
		}
		tempRow := &pb.ReportSalesHeader{
			PrintDate:        row.PrintDate,
			ParamDate:        row.ParamDate,
			TotalOmset:       row.TotalOmset,
			TotalLabaGross:   row.TotalLabaGross,
			TotalPenjualan:   row.TotalPenjualan,
			TotalBarang:      row.TotalBarang,
			ReportSalesLines: sliceLines,
		}
		temp = append(temp, tempRow)
	}
	return &pb.ResReportSalesHeader{
		Total: data.Total,
		Data:  temp,
	}, nil
}
