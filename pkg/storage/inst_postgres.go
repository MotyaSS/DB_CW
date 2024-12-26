package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type InstPostgres struct {
	db *sqlx.DB
}

const (
	pageSize = 20
)

func newInstPostgres(db *sqlx.DB) *InstPostgres {
	return &InstPostgres{db: db}
}

func (s *InstPostgres) GetInstrument(id int) (entity.Instrument, error) {
	var instrument entity.Instrument
	query := fmt.Sprintf("SELECT * FROM %s WHERE instrument_id = $1", instrumentsTable)
	err := s.db.Get(&instrument, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return instrument, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("instrument with id %d not found", id),
		}
	}
	return instrument, err
}

func (s *InstPostgres) GetAllInstruments(filter entity.InstFilter) ([]entity.Instrument, error) {
	instruments := make([]entity.Instrument, 0)

	query := fmt.Sprintf(`
		SELECT i.* FROM %s i
		LEFT JOIN categories c ON i.category_id = c.category_id
		LEFT JOIN manufacturers m ON i.manufacturer_id = m.manufacturer_id
		WHERE TRUE`, instrumentsTable)

	args := make([]interface{}, 0)
	argId := 1

	if filter.Category != nil {
		query += fmt.Sprintf(" AND c.category_name = $%d", argId)
		args = append(args, *filter.Category)
		argId++
	}

	if filter.Manufacturer != nil {
		query += fmt.Sprintf(" AND m.manufacturer_name = $%d", argId)
		args = append(args, *filter.Manufacturer)
		argId++
	}

	if filter.PriceFloor != nil {
		query += fmt.Sprintf(" AND i.price_per_day >= $%d", argId)
		args = append(args, *filter.PriceFloor)
		argId++
	}

	if filter.PriceCeil != nil {
		query += fmt.Sprintf(" AND i.price_per_day <= $%d", argId)
		args = append(args, *filter.PriceCeil)
		argId++
	}

	// Add pagination
	offset := (filter.Page - 1) * pageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	err := s.db.Select(&instruments, query, args...)
	return instruments, err
}

func (s *InstPostgres) CreateInstrument(instrument entity.Instrument) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO %s (instrument_name, category_id, store_id, manufacturer_id, description, price_per_day) VALUES ($1, $2, $3, $4, $5, $6) RETURNING instrument_id", instrumentsTable)
	err = s.db.QueryRow(query,
		instrument.InstrumentName,
		instrument.CategoryId,
		instrument.StoreId,
		instrument.ManufacturerId,
		instrument.Description,
		instrument.PricePerDay).Scan(&id)

	if err == nil {
		return id, err
	}

	var pqErr *pq.Error
	ok := errors.As(err, &pqErr)
	if !ok {
		return id, err
	}

	switch pqErr.Code.Name() {
	case "foreign_key_violation":
		var field string
		var value int
		switch {
		case strings.Contains(pqErr.Detail, "category_id"):
			field = "category"
			value = instrument.CategoryId
			break
		case strings.Contains(pqErr.Detail, "store_id"):
			field = "store"
			value = instrument.StoreId
			break
		case strings.Contains(pqErr.Detail, "manufacturer_id"):
			field = "manufacturer"
			value = instrument.ManufacturerId
			break
		}
		return id, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("%s with id %d not found", field, value),
		}
	}

	slog.Error("unknown internal server error during creating instrument",
		"err", pqErr.Message,
	)

	return id, &httpError.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        fmt.Sprintf("internal server error during creating instrument"),
	}
}
