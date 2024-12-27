package service

import (
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

	if userRole.RoleId < entity.RoleStaffId {
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
