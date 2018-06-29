package main

import (
	"context"
	"fmt"
	"gokit-grpc/car-microservice/car"
	"gokit-grpc/car-microservice/database"
	"gokit-grpc/car-microservice/pb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	carService := &car.CarService{database.NewDBCarRepository()}
	errors := make(chan error)

	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			errors <- err
			return
		}

		gRPCServer := grpc.NewServer()
		pb.RegisterCarServer(gRPCServer, car.NewGRPCServer(ctx, car.Endpoints{
			GetCarEndpoint:    car.MakeGetCarEndpoint(carService),
			CreateCarEndpoint: car.MakeCreateCarEndpoint(carService),
		}))

		fmt.Println("gRPC listen on 9090")
		errors <- gRPCServer.Serve(listener)
	}()

	fmt.Println(<-errors)
}
