package storage

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type InstPostgres struct {
	db *sqlx.DB
}

const (
	pageSize = 20
)

func newInstPostgres(db *sqlx.DB) *InstPostgres {
	return &InstPostgres{db: db}
}

func (s *InstPostgres) GetInstrument(id int) (entity.Instrument, error) {
	return entity.Instrument{}, nil
}

func (s *InstPostgres) GetAllInstruments(filter entity.InstFilter) ([]entity.Instrument, error) {
	return nil, nil
}

func (s *InstPostgres) CreateInstrument([]entity.Instrument) (id int, err error) {
	return 0, nil
}
