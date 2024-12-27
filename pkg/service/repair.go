package service

import (
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type RepairService struct {
	storage storage.Repair
	auth    Authorisation
}

func NewRepairService(storage storage.Repair, auth Authorisation) *RepairService {
	return &RepairService{storage: storage, auth: auth}
}

func (s *RepairService) GetRepair(callerId int, id int) (entity.Repair, error) {
	role, err := s.auth.GetUserRole(callerId)
	if err != nil {
		return entity.Repair{}, err
	}
	if (!s.auth.HasPermission(role, entity.Role{RoleId: entity.RoleStaff.RoleId})) {
		return entity.Repair{}, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "user has no permission to view repairs",
		}

	}
	return s.storage.GetRepair(id)
}
func (s *RepairService) CreateRepair(callerId int, repair entity.Repair) (id int, err error) {
	return s.storage.CreateRepair(callerId, repair)
}
func (s *RepairService) GetInstrumentRepairs(callerId int, instrumentId int) ([]entity.Repair, error) {
	return s.storage.GetInstrumentRepairs(instrumentId)
}
