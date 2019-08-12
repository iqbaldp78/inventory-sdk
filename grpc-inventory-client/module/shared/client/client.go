package client

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/module/shared/client/pb"
)

//Client used as GRPC client class
type Client struct {
	conn    *grpc.ClientConn
	service pb.InventoryServiceClient
}

//NewClient used as GRPC client contructor
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewInventoryServiceClient(conn)
	return &Client{conn, c}, nil
}

//Close used for closed client connection
func (c *Client) Close() {
	c.conn.Close()
}

//GetOnhandQTY used for call GetOnhandQTY function on server
func (c *Client) GetOnhandQTY(ctx context.Context) (*pb.ResOnhandQTY, error) {
	return c.service.GetOnhandQTY(ctx, &pb.Empty{})
}

//GetInbounds used for call GetInbounds function on server
func (c *Client) GetInbounds(ctx context.Context) (*pb.ResInbounds, error) {
	return c.service.GetInbounds(ctx, &pb.Empty{})
}

//GetOutbounds used for call GetInbounds function on server
func (c *Client) GetOutbounds(ctx context.Context) (*pb.ResOutbounds, error) {
	return c.service.GetOutbounds(ctx, &pb.Empty{})
}

//GetReportValue used for call GetReportValue function on server
func (c *Client) GetReportValue(ctx context.Context) (*pb.ResReportValue, error) {
	return c.service.GetReportValue(ctx, &pb.Empty{})
}

//GetReportSales used for call GetReportSales function on server
func (c *Client) GetReportSales(ctx context.Context, param *pb.ParamDate) (*pb.ResReportSalesHeader, error) {
	return c.service.GetReportSales(ctx, param)
}
