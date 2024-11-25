package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable         = "users"
	rolesTable         = "roles"
	storesTable        = "stores"
	categoriesTable    = "categories"
	manufacturersTable = "manufacturers"
	instrumentsTable   = "instruments"
	rentalsTable       = "rentals"
	paymentsTable      = "payments"
	repairsTable       = "repairs"
	discountsTable     = "repairs"
	reviewsTable       = "reviews"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
