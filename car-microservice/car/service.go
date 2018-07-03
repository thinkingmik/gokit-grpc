package car

import (
	"gokit-grpc/car-microservice/datastore"
	"gokit-grpc/car-microservice/models"
)

type IService interface {
	GetCar(id int32) (*models.Car, error)
	CreateCar(name string, manifacturer string) (*models.Car, error)
}

type CarService struct {
	Repository datastore.IDataStore
}

// Business logic here

func (s *CarService) GetCar(id int32) (*models.Car, error) {
	return s.Repository.FindById(id)
}

func (s *CarService) CreateCar(name string, manifacturer string) (*models.Car, error) {
	return s.Repository.Create(name, manifacturer)
}
