package service

import (
	"log/slog"
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type InstService struct {
	storage storage.Instrument
	auth    Authorisation
}

func NewInstService(storage storage.Instrument, auth Authorisation) *InstService {
	return &InstService{
		storage: storage,
		auth:    auth,
	}
}

func (s *InstService) GetAllInstruments(filter entity.InstFilter) ([]InstrumentWithDiscount, error) {
	instruments, err := s.storage.GetAllInstruments(filter)
	if err != nil {
		return nil, err
	}
	res := make([]InstrumentWithDiscount, 0, len(instruments))
	for _, inst := range instruments {
		discount, err := s.GetActiveDiscount(inst.InstrumentId)
		if err != nil {
			return nil, err
		}
		res = append(res, InstrumentWithDiscount{inst, discount})
	}
	return res, nil
}

func (s *InstService) GetInstrument(id int) (InstrumentWithDiscount, error) {
	inst, err := s.storage.GetInstrument(id)
	if err != nil {
		return InstrumentWithDiscount{}, err
	}
	discount, err := s.GetActiveDiscount(id)
	if err != nil {
		return InstrumentWithDiscount{}, err
	}
	return InstrumentWithDiscount{inst, discount}, nil
}
func (s *InstService) CreateInstrument(instrument entity.Instrument) (id int, err error) {
	return s.storage.CreateInstrument(instrument)
}

func (s *InstService) GetActiveDiscount(instrumentId int) (*entity.Discount, error) {
	return s.storage.GetActiveDiscount(instrumentId)
}

func (s *InstService) DeleteInstrument(callerId int, instrumentId int) error {
	// Check if user has staff or higher permissions
	userRole, err := s.auth.GetUserRole(callerId)
	if err != nil {
		return err
	}

	if !s.auth.HasPermission(userRole, entity.Role{RoleId: entity.RoleStaff.RoleId}) {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "insufficient  permissions to delete instruments",
		}
	}

	// Check if instrument exists
	_, err = s.storage.GetInstrument(instrumentId)
	if err != nil {
		return err
	}

	// Delete the instrument
	return s.storage.DeleteInstrument(instrumentId)
}

func (s *InstService) GetCategories() ([]entity.Category, error) {
	return s.storage.GetCategories()
}

func (s *InstService) GetManufacturers() ([]entity.Manufacturer, error) {
	return s.storage.GetManufacturers()
}

func (s *InstService) CreateCategory(callerId int, category entity.Category) (int, error) {
	user, err := s.auth.GetUserById(callerId)
	if err != nil {
		return 0, err
	}
	slog.Info("user", "user", user)
	userRole, err := s.auth.GetRole(user.RoleId)
	slog.Info("userRole", "userRole", userRole)	
	if err != nil {
		return 0, err
	}

	if !s.auth.HasPermission(userRole, entity.RoleChief) {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "недостаточно прав для создания категорий",
		}
	}

	return s.storage.CreateCategory(category)
}

func (s *InstService) CreateManufacturer(callerId int, manufacturer entity.Manufacturer) (int, error) {
	// Проверяем права (должен быть Chief или выше)
	userRole, err := s.auth.GetUserRole(callerId)
	if err != nil {
		return 0, err
	}

	if !s.auth.HasPermission(userRole, entity.RoleChief) {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "недостаточно прав для создания производителей",
		}
	}

	return s.storage.CreateManufacturer(manufacturer)
}
