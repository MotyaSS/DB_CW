package service

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type InstService struct {
	storage storage.Instrument
}

func NewInstService(storage storage.Instrument) *InstService {
	return &InstService{storage: storage}
}

func (s *InstService) GetAllInstruments(filter entity.InstFilter) ([]entity.Instrument, error) {
	return s.storage.GetAllInstruments(filter)
}
