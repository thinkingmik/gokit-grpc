package car

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type CarResponse struct {
	ID           int32
	Name         string
	Manifacturer string
}

type GetCarRequest struct {
	ID int32
}

type CreateCarRequest struct {
	Name         string
	Manifacturer string
}

type Endpoints struct {
	GetCarEndpoint    endpoint.Endpoint
	CreateCarEndpoint endpoint.Endpoint
}

// Validation logic here

func MakeGetCarEndpoint(svc IService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(GetCarRequest)
		car, _ := svc.GetCar(request.ID)
		if car == nil {
			return CarResponse{}, nil
		}
		return CarResponse{car.ID, car.Name, car.Manifacturer}, nil
	}
}

func MakeCreateCarEndpoint(svc IService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(CreateCarRequest)
		car, _ := svc.CreateCar(request.Name, request.Manifacturer)
		return CarResponse{
			car.ID,
			car.Name,
			car.Manifacturer,
		}, nil
	}
}
