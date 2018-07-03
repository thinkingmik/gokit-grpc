package datastore

import "gokit-grpc/car-microservice/models"

type MemoryDataStore struct {
	cursor int32
	List   map[int32]*models.Car
}

func NewMemoryDataStore(conf map[string]string) (IDataStore, error) {
	return &MemoryDataStore{
		0,
		make(map[int32]*models.Car),
	}, nil
}

func (ds *MemoryDataStore) FindById(id int32) (*models.Car, error) {
	return ds.List[id], nil
}

func (ds *MemoryDataStore) Create(name string, manifacturer string) (*models.Car, error) {
	ds.cursor++
	car := &models.Car{
		ds.cursor,
		name,
		manifacturer,
	}
	ds.List[car.ID] = car
	return car, nil
}
