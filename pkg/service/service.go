package service

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type Authorisation interface {
	GetAllRoles() ([]entity.Role, error)
	GetUserRole(userId int) (entity.Role, error)

	CreateUser(callerId int, user entity.User) (int, error)
	CreateCustomer(user entity.User) (int, error)
	CreateStaff(callerId int, user entity.User) (int, error)
	CreateChief(callerId int, user entity.User) (int, error)
	CreateAdmin(callerId int, user entity.User) (int, error)

	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Instrument interface {
	GetAllInstruments(filter entity.InstFilter) ([]entity.Instrument, error)
	GetInstrument(id int) (entity.Instrument, error)
	CreateInstrument(instrument entity.Instrument) (id int, err error)
}

type Review interface {
}

type Rent interface {
}

type Store interface {
}

type Service struct {
	Authorisation
	Instrument
	Review
	Rent
	Store
}

func New(storage *storage.Storage) *Service {
	return &Service{
		Authorisation: NewAuthService(storage),
		Instrument:    NewInstService(storage),
	}
}
