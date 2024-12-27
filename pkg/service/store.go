package service

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type StoreService struct {
	storage storage.Store
}

func NewStoreService(storage storage.Store) *StoreService {
	return &StoreService{storage: storage}
}

func (s *StoreService) GetAllStores() ([]entity.Store, error) {
	return s.storage.GetAllStores()
}

func (s *StoreService) GetStore(id int) (entity.Store, error) {
	return s.storage.GetStore(id)
}

func (s *StoreService) CreateStore(store entity.Store) (int, error) {
	return s.storage.CreateStore(store)
}

func (s *StoreService) DeleteStore(id int) error {
	return s.storage.DeleteStore(id)
}
