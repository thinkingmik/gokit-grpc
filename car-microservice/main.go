package main

import (
	"context"
	"fmt"
	"gokit-grpc/car-microservice/car"
	"gokit-grpc/car-microservice/datastore"
	"gokit-grpc/car-microservice/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ds, err := datastore.CreateDatastore(map[string]string{
		"DATASTORE": "mysql",
		"DSN":       "root:root@tcp(127.0.0.1:3306)/golang",
	})
	if err != nil {
		log.Panicf("Failed to create datastore", err.Error())
	}
	carService := &car.CarService{ds}
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
