package database

type Car struct {
	ID           int32
	Name         string
	Manifacturer string
}

type CarRepository interface {
	FindById(id int32) (*Car, error)
	Create(name string, manifacturer string) (*Car, error)
}

type DBCarRepository struct {
	cursor int32
	DB     map[int32]*Car
}

func (r *DBCarRepository) FindById(id int32) (*Car, error) {
	return r.DB[id], nil
}

func (r *DBCarRepository) Create(name string, manifacturer string) (*Car, error) {
	r.cursor++
	car := &Car{
		r.cursor,
		name,
		manifacturer,
	}
	r.DB[car.ID] = car
	return car, nil
}

func NewDBCarRepository() *DBCarRepository {
	return &DBCarRepository{
		0,
		make(map[int32]*Car),
	}
}
