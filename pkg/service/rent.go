package service

import (
	"time"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type RentService struct {
	storage storage.Rent
}

func NewRentService(storage storage.Rent) *RentService {
	return &RentService{storage: storage}
}

func (s *RentService) CreateRental(userId int, instrumentId int) (int, error) {
	// Create a new rental with current time
	rental := entity.Rental{
		UserId:       userId,
		InstrumentId: instrumentId,
		RentalDate:   time.Now(),
	}

	return s.storage.CreateRental(rental)
}

func (s *RentService) GetRental(id int) (entity.Rental, error) {
	return s.storage.GetRental(id)
}

func (s *RentService) GetUserRentals(userId int) ([]entity.Rental, error) {
	return s.storage.GetUserRentals(userId)
}

func (s *RentService) GetInstrumentRentals(instrumentId int) ([]entity.Rental, error) {
	return s.storage.GetInstrumentRentals(instrumentId)
}

func (s *RentService) DeleteRental(id int) error {
	return s.storage.DeleteRental(id)
}

func (s *RentService) ReturnInstrument(rentalId int) error {
	return s.storage.ReturnInstrument(rentalId)
}
