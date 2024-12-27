package service

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type Authorisation interface {
	GetAllRoles() ([]entity.Role, error)
	GetUserRole(userId int) (entity.Role, error)
	CreateUser(callerId int, user entity.User) (int, error)
	CreateCustomer(user entity.User) (int, error)
	HasPermission(userRole, requiredRole entity.Role) bool
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type InstrumentWithDiscount struct {
	Instrument entity.Instrument
	Discount   *entity.Discount
}

type Instrument interface {
	GetAllInstruments(filter entity.InstFilter) ([]InstrumentWithDiscount, error)
	GetInstrument(id int) (InstrumentWithDiscount, error)
	CreateInstrument(instrument entity.Instrument) (id int, err error)
	GetActiveDiscount(instrumentId int) (*entity.Discount, error)
	DeleteInstrument(callerId int, instrumentId int) error
}

type Repair interface {
	GetRepair(callerId int, id int) (entity.Repair, error)
	CreateRepair(callerId int, repair entity.Repair) (id int, err error)
	GetInstrumentRepairs(callerId int, instrumentId int) ([]entity.Repair, error)
}

type Review interface {
}

type Rent interface {
	CreateRental(userId int, instrumentId int) (int, error)
	GetRental(id int) (entity.Rental, error)
	GetUserRentals(userId int) ([]entity.Rental, error)
	GetInstrumentRentals(instrumentId int) ([]entity.Rental, error)
	DeleteRental(id int) error
	ReturnInstrument(rentalId int) error
}

type Store interface {
	GetAllStores() ([]entity.Store, error)
	GetStore(id int) (entity.Store, error)
	CreateStore(store entity.Store) (int, error)
	DeleteStore(id int) error
}

type Service struct {
	Authorisation
	Instrument
	Review
	Rent
	Store
	Repair
}

func New(storage *storage.Storage) *Service {
	auth := NewAuthService(storage)
	return &Service{
		Authorisation: auth,
		Instrument:    NewInstService(storage, auth),
		Store:         NewStoreService(storage),
		Rent:          NewRentService(storage),
	}
}
