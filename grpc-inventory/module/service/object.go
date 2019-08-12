package service

import (
	"context"

	"github.com/gocraft/dbr"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/lib/sqlite"
)

//Service _
type Service interface {
	GetOnhandQTY(ctx context.Context, empty *pb.Empty) (*pb.ResOnhandQTY, error)
	GetInbounds(ctx context.Context, empty *pb.Empty) (*pb.ResInbounds, error)
	GetOutbounds(ctx context.Context, empty *pb.Empty) (*pb.ResOutbounds, error)
	GetReportValue(ctx context.Context, empty *pb.Empty) (*pb.ResReportValue, error)
	GetReportSales(ctx context.Context, param *pb.ParamDate) (*pb.ResReportSalesHeader, error)
}

type service struct {
	db *dbr.Session
}

//New _
func New() Service {
	return &service{sqlite.GetDB()}
}
