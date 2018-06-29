package car

import (
	"gokit-grpc/car-microservice/database"
)

type Service interface {
	GetCar(id int32) (*database.Car, error)
	CreateCar(name string, manifacturer string) (*database.Car, error)
}

type CarService struct {
	Repository database.CarRepository
}

func (s *CarService) GetCar(id int32) (*database.Car, error) {
	return s.Repository.FindById(id)
}

func (s *CarService) CreateCar(name string, manifacturer string) (*database.Car, error) {
	return s.Repository.Create(name, manifacturer)
}
