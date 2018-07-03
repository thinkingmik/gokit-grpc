package models

type DataStoreConnectionError struct {
	Message string
}

func (e *DataStoreConnectionError) Error() string { return e.Message }
