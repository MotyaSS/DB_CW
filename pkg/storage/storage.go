package storage

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
	GetRole(roleId int) (entity.Role, error)
}

type Instrument interface {
}

type Review interface {
}

type Rent interface {
}

type Store interface {
}

type Storage struct {
	Authorisation
	Instrument
	Review
	Rent
	Store
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		Authorisation: NewAuthPostgres(db),
	}
}
