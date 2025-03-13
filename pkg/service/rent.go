package service

import (
	"net/http"
	"time"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type RentService struct {
	storage *storage.Storage
}

func NewRentService(storage *storage.Storage) *RentService {
	return &RentService{
		storage: storage,
	}
}

func (s *RentService) CreateRental(userId, instrumentId int, startDateStr, endDateStr string) (int, error) {
	// Парсим даты
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusBadRequest,
			Msg:        "неверный формат даты начала аренды",
		}
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusBadRequest,
			Msg:        "неверный формат даты окончания аренды",
		}
	}

	// Проверяем корректность дат
	if startDate.After(endDate) {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusBadRequest,
			Msg:        "дата начала аренды не может быть позже даты окончания",
		}
	}

	if startDate.Before(time.Now()) {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusBadRequest,
			Msg:        "дата начала аренды не может быть в прошлом",
		}
	}

	// Проверяем, не занят ли инструмент на эти даты
	existingRentals, err := s.storage.GetInstrumentRentals(instrumentId)
	if err != nil {
		return 0, err
	}

	for _, rental := range existingRentals {
		// Проверяем пересечение периодов
		if (startDate.Before(rental.ReturnDate) && endDate.After(rental.RentalDate)) ||
			(startDate.Equal(rental.RentalDate) || endDate.Equal(rental.ReturnDate)) {
			return 0, &httpError.ErrorWithStatusCode{
				HTTPStatus: http.StatusConflict,
				Msg:        "инструмент уже арендован на эти даты",
			}
		}
	}

	// Создаем запись об аренде
	rental := entity.Rental{
		UserId:       userId,
		InstrumentId: instrumentId,
		RentalDate:   startDate,
		ReturnDate:   endDate,
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
