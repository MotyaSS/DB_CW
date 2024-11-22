package storage

import (
	"github.com/jmoiron/sqlx"
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

type Storage struct {
	Authorisation
	User
	Instrument
	Review
	Rent
	Store
}

func New(db *sqlx.DB) *Storage {
	return &Storage{}
}
