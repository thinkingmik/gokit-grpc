package car

import (
	"context"

	"gokit-grpc/car-microservice/pb"
)

func EncodeCarResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(CarResponse)
	return &pb.CarResponse{
		Id:           res.ID,
		Name:         res.Name,
		Manifacturer: res.Manifacturer,
	}, nil
}

func DecodeGetCarRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetCarRequest)
	return GetCarRequest{
		ID: req.Id,
	}, nil
}

func DecodeCreateCarRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateCarRequest)
	return CreateCarRequest{
		Name:         req.Name,
		Manifacturer: req.Manifacturer,
	}, nil
}
