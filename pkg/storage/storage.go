package storage

import (
	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
	GetRole(roleId int) (entity.Role, error)
	GetRoleId(roleName string) (int, error)
	GetAllRoles() ([]entity.Role, error)
	GetUserRole(userId int) (entity.Role, error)
	GetAllUsers() ([]entity.User, error)
	GetUserById(userId int) (entity.User, error)
	DeleteUser(userId int) error
}

type Instrument interface {
	GetInstrument(id int) (entity.Instrument, error)
	GetAllInstruments(filter entity.InstFilter) ([]entity.Instrument, error)
	CreateInstrument(entity.Instrument) (id int, err error)
	GetActiveDiscount(instrumentId int) (*entity.Discount, error)
	DeleteInstrument(id int) error
	GetCategories() ([]entity.Category, error)
	GetManufacturers() ([]entity.Manufacturer, error)
	CreateCategory(category entity.Category) (int, error)
	CreateManufacturer(manufacturer entity.Manufacturer) (int, error)
}

type Review interface {
	GetAllReviews(instrumentId int) ([]entity.Review, error)
	GetReview(id int) (entity.Review, error)
	CreateReview(callerId int, review entity.Review) (int, error)
	DeleteReview(callerId int, reviewId int) error
}

type Rent interface {
	CreateRental(rental entity.Rental) (int, error)
	GetRental(id int) (entity.Rental, error)
	GetUserRentals(userId int) ([]entity.Rental, error)
	GetInstrumentRentals(instrumentId int) ([]entity.Rental, error)
	UpdateRental(rental entity.Rental) error
	DeleteRental(id int) error
	ReturnInstrument(rentalId int) error
}

type Store interface {
	GetAllStores() ([]entity.Store, error)
	GetStore(id int) (entity.Store, error)
	CreateStore(store entity.Store) (int, error)
	DeleteStore(id int) error
}

type Storage struct {
	Authorisation
	Instrument
	Review
	Rent
	Store
	Repair
}

type Repair interface {
	GetRepair(id int) (entity.Repair, error)
	CreateRepair(callerId int, repair entity.Repair) (id int, err error)
	GetInstrumentRepairs(instrumentId int) ([]entity.Repair, error)
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		Authorisation: NewAuthPostgres(db),
		Instrument:    newInstPostgres(db),
		Rent:          newRentPostgres(db),
		Store:         newStorePostgres(db),
		Repair:        NewRepairPostgres(db),
		Review:        newReviewPostgres(db),
	}
}
