package module

import (
	"context"
	"fmt"
	"net"

	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service service.Service
}

//ListenGRPC used for run grpc service
func ListenGRPC(s service.Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterInventoryServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) GetOnhandQTY(ctx context.Context, empty *pb.Empty) (*pb.ResOnhandQTY, error) {
	return s.service.GetOnhandQTY(ctx, empty)
}

func (s *grpcServer) GetInbounds(ctx context.Context, empty *pb.Empty) (*pb.ResInbounds, error) {
	return s.service.GetInbounds(ctx, empty)
}

func (s *grpcServer) GetOutbounds(ctx context.Context, empty *pb.Empty) (*pb.ResOutbounds, error) {
	return s.service.GetOutbounds(ctx, empty)
}

func (s *grpcServer) GetReportValue(ctx context.Context, empty *pb.Empty) (*pb.ResReportValue, error) {
	return s.service.GetReportValue(ctx, empty)
}

func (s *grpcServer) GetReportSales(ctx context.Context, param *pb.ParamDate) (*pb.ResReportSalesHeader, error) {
	return s.service.GetReportSales(ctx, param)
}
