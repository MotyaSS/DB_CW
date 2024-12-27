package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/jmoiron/sqlx"
)

type StorePostgres struct {
	db *sqlx.DB
}

func newStorePostgres(db *sqlx.DB) *StorePostgres {
	return &StorePostgres{db: db}
}

func (s *StorePostgres) GetAllStores() ([]entity.Store, error) {
	var stores []entity.Store
	query := fmt.Sprintf(`SELECT * FROM %s`, storesTable)
	err := s.db.Select(&stores, query)
	if errors.Is(err, sql.ErrNoRows) {
		return stores, nil
	}

	if err != nil {
		slog.Error("error getting all stores", "err", err.Error())
		return nil, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "internal server error during getting stores",
		}
	}
	return stores, nil
}

func (s *StorePostgres) GetStore(id int) (entity.Store, error) {
	var store entity.Store
	query := fmt.Sprintf(`SELECT * FROM %s WHERE store_id = $1`, storesTable)
	err := s.db.Get(&store, query, id)

	if err == nil {
		return store, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return store, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("store with id %d not found", id),
		}
	}

	slog.Error("unknown error getting store", "err", err.Error())
	return store, &httpError.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "internal server error during getting store",
	}
}

func (s *StorePostgres) CreateStore(store entity.Store) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (store_name, address, phone_number) 
		VALUES ($1, $2, $3) RETURNING store_id`, storesTable)

	err := s.db.QueryRow(query,
		store.StoreName,
		store.StoreAddress,
		store.PhoneNumber).Scan(&id)

	if err != nil {
		slog.Error("error creating store", "err", err.Error())
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "internal server error during creating store",
		}
	}

	return id, nil
}

func (s *StorePostgres) DeleteStore(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE store_id = $1`, storesTable)
	result, err := s.db.Exec(query, id)

	if err != nil {
		slog.Error("error deleting store", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "internal server error during deleting store",
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("error getting rows affected", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "internal server error during deleting store",
		}
	}

	if rowsAffected == 0 {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("store with id %d not found", id),
		}
	}

	return nil
}
