package datastore

import "gokit-grpc/car-microservice/models"

//var UserNotFoundError = errors.New("User not found")

type IDataStore interface {
	FindById(id int32) (*models.Car, error)
	Create(name string, manifacturer string) (*models.Car, error)
}

func CreateDatastore(conf map[string]string) (IDataStore, error) {
	engineName, def := conf["DATASTORE"]
	if !def {
		engineName = "memory"
	}

	if engineName == "mysql" {
		return NewMySQLDataStore(conf)
	} else {
		return NewMemoryDataStore(conf)
	}
}
