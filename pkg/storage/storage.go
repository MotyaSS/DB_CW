package storage

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
	GetRole(roleId int) (entity.Role, error)
	GetRoleId(roleName string) (int, error)
	GetAllRoles() ([]entity.Role, error)
}

type Instrument interface {
	GetInstrument(id int) (entity.Instrument, error)
	GetAllInstruments(filter entity.InstFilter) ([]entity.Instrument, error)
	CreateInstrument([]entity.Instrument) (id int, err error)
}

type Review interface {
	GetAllReviews() ([]entity.Review, error)
	GetReview(id int) (entity.Review, error)
	CreateReview(callerId int, review []entity.Review) error
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
		Instrument:    newInstPostgres(db),
	}
}
