package service

import (
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type Authorisation interface {
}

type User interface {
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
	User
	Instrument
	Review
	Rent
	Store
	storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}
