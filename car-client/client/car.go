package client

import (
	"context"
	"gokit-grpc/car-client/pb"
	"log"
)

type ICarClient interface {
	GetCar(id int32) *pb.CarResponse
	CreateCar(name string, manifacturer string) *pb.CarResponse
}

type CarGRPCClient struct {
	Client *GRPCClient
}

func (c *CarGRPCClient) GetCar(id int32) *pb.CarResponse {
	ctx := context.Background()
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := pb.NewCarClient(conn)
	res, err := client.GetCar(ctx, &pb.GetCarRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func (c *CarGRPCClient) CreateCar(name string, manifacturer string) *pb.CarResponse {
	ctx := context.Background()
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := pb.NewCarClient(conn)
	res, err := client.CreateCar(ctx, &pb.CreateCarRequest{
		Name:         name,
		Manifacturer: manifacturer,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func NewCarGRPCClient(host string) ICarClient {
	return &CarGRPCClient{NewGRPCClient(host)}
}
