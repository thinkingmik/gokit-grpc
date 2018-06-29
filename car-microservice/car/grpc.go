package car

import (
	"context"
	"gokit-grpc/car-microservice/pb"

	"github.com/go-kit/kit/transport/grpc"
)

type GRPCServer struct {
	getCar    grpc.Handler
	createCar grpc.Handler
}

func (s *GRPCServer) GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.CarResponse, error) {
	_, resp, err := s.getCar.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CarResponse), nil
}

func (s *GRPCServer) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CarResponse, error) {
	_, resp, err := s.createCar.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CarResponse), nil
}

func NewGRPCServer(_ context.Context, endpoint Endpoints) pb.CarServer {
	return &GRPCServer{
		getCar: grpc.NewServer(
			endpoint.GetCarEndpoint,
			DecodeGetCarRequest,
			EncodeCarResponse,
		),
		createCar: grpc.NewServer(
			endpoint.CreateCarEndpoint,
			DecodeCreateCarRequest,
			EncodeCarResponse,
		),
	}
}
