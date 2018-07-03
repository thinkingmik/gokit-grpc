package datastore

import (
	"database/sql"
	"errors"
	"fmt"
	"gokit-grpc/car-microservice/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDataStore struct {
	Dsn       string
	Connector *sql.DB
}

func NewMySQLDataStore(conf map[string]string) (IDataStore, error) {
	dsn, ok := conf["DSN"]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s is required for the mysql datastore", "DSN"))
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("Failed to connect to datastore: %s", err.Error())
		return nil, &models.DataStoreConnectionError{
			"Failed to connect to datastore",
		}
	}

	return &MySQLDataStore{
		dsn,
		db,
	}, nil
}

func (ds *MySQLDataStore) FindById(key int32) (*models.Car, error) {
	var id int32
	var name string
	var manifacturer string

	sqlStatement := "SELECT id, name, manifacturer FROM cars WHERE id=?;"
	row := ds.Connector.QueryRow(sqlStatement, key)

	switch err := row.Scan(&id, &name, &manifacturer); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &models.Car{
			id,
			name,
			manifacturer,
		}, nil
	default:
		return nil, err
	}
}

func (ds *MySQLDataStore) Create(name string, manifacturer string) (*models.Car, error) {
	return nil, nil
}
