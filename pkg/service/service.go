package service

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type Authorisation interface {
	CreateUser(callerJWT string, user entity.User) (int, error)
	CreateCustomer(user entity.User) (int, error)
	CreateStaff(callerId int, user entity.User) (int, error)
	CreateChief(callerId int, user entity.User) (int, error)
	CreateAdmin(callerId int, user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) int
}

type Instrument interface {
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
	}
}
