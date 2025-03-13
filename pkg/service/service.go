package service

import (
	"github.com/MotyaSS/DB_CW/pkg/config"
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type Authorisation interface {
	GetAllRoles() ([]entity.Role, error)
	GetRole(roleId int) (entity.Role, error)
	GetUserRole(userId int) (entity.Role, error)
	CreateUser(callerId int, user entity.User) (int, error)
	CreateCustomer(user entity.User) (int, error)
	HasPermission(userRole, requiredRole entity.Role) bool
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CheckPermission(userId int, requiredRole entity.Role) error
	GetAllUsers() ([]entity.User, error)
	GetUserById(userId int) (entity.User, error)
	DeleteUser(userId int) error
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
	GetCategories() ([]entity.Category, error)
	GetManufacturers() ([]entity.Manufacturer, error)
	CreateCategory(callerId int, category entity.Category) (int, error)
	CreateManufacturer(callerId int, manufacturer entity.Manufacturer) (int, error)
}

type Repair interface {
	GetRepair(callerId int, id int) (entity.Repair, error)
	CreateRepair(callerId int, repair entity.Repair) (id int, err error)
	GetInstrumentRepairs(callerId int, instrumentId int) ([]entity.Repair, error)
}

type Review interface {
}

type Rent interface {
	CreateRental(userId, instrumentId int, startDate, endDate string) (int, error)
	GetRental(rentalId int) (entity.Rental, error)
	GetUserRentals(userId int) ([]entity.Rental, error)
	GetInstrumentRentals(instrumentId int) ([]entity.Rental, error)
	ReturnInstrument(rentalId int) error
}

type Store interface {
	GetAllStores() ([]entity.Store, error)
	GetStore(id int) (entity.Store, error)
	CreateStore(store entity.Store) (int, error)
	DeleteStore(id int) error
}

type Backup interface {
	CreateBackup() (string, error)
	RestoreFromBackup(backupName string) error
	ListBackups() ([]string, error)
}

type Service struct {
	Authorisation
	Instrument
	Review
	Rent
	Store
	Repair
	Backup
}

func New(storage *storage.Storage, config config.Database) *Service {
	auth := NewAuthService(storage)
	backupService := NewBackupService(config)

	return &Service{
		Authorisation: auth,
		Instrument:    NewInstService(storage, auth),
		Store:         NewStoreService(storage),
		Rent:          NewRentService(storage),
		Repair:        NewRepairService(storage, auth),
		Backup:        backupService,
	}
}
