package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

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
	query := fmt.Sprintf(`
		SELECT instrument_id, instrument_name, category_id, 
			   manufacturer_id, store_id, description, 
			   price_per_day, image_url 
		FROM %s WHERE instrument_id = $1`, instrumentsTable)
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
		SELECT i.instrument_id, i.instrument_name, i.category_id, 
			   i.manufacturer_id, i.store_id, i.description, 
			   i.price_per_day, i.image_url 
		FROM %s i
		LEFT JOIN categories c ON i.category_id = c.category_id
		LEFT JOIN manufacturers m ON i.manufacturer_id = m.manufacturer_id
		WHERE TRUE`, instrumentsTable)

	args := make([]interface{}, 0)
	argId := 1

	if len(filter.Categories) > 0 {
		placeholders := make([]string, len(filter.Categories))
		for i, category := range filter.Categories {
			placeholders[i] = fmt.Sprintf("$%d", argId)
			args = append(args, category)
			argId++
		}
		query += fmt.Sprintf(" AND c.category_name = ANY(ARRAY[%s])", strings.Join(placeholders, ", "))
	}

	if len(filter.Manufacturers) > 0 {
		placeholders := make([]string, len(filter.Manufacturers))
		for i, manufacturer := range filter.Manufacturers {
			placeholders[i] = fmt.Sprintf("$%d", argId)
			args = append(args, manufacturer)
			argId++
		}
		query += fmt.Sprintf(" AND m.manufacturer_name = ANY(ARRAY[%s])", strings.Join(placeholders, ", "))
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
	query := fmt.Sprintf(`
		INSERT INTO %s (
			instrument_name, category_id, store_id, 
			manufacturer_id, description, price_per_day, 
			image_url
		) VALUES ($1, $2, $3, $4, NULLIF($5, ''), $6, $7) 
		RETURNING instrument_id
	`, instrumentsTable)

	err = s.db.QueryRow(
		query,
		instrument.InstrumentName,
		instrument.CategoryId,
		instrument.StoreId,
		instrument.ManufacturerId,
		instrument.Description,
		instrument.PricePerDay,
		instrument.ImageUrl,
	).Scan(&id)

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
		case strings.Contains(pqErr.Detail, "store_id"):
			field = "store"
			value = instrument.StoreId
		case strings.Contains(pqErr.Detail, "manufacturer_id"):
			field = "manufacturer"
			value = instrument.ManufacturerId
		}
		return id, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("%s with id %d not found", field, value),
		}
	}

	slog.Error("unknown internal server error during creating instrument",
		"err", pqErr.Message)

	return id, &httpError.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "internal server error during creating instrument",
	}
}

func (s *InstPostgres) GetActiveDiscount(instrumentId int) (*entity.Discount, error) {
	var discount entity.Discount
	query := fmt.Sprintf(`
		SELECT discount_id, instrument_id, discount_percentage, valid_until 
		FROM %s 
		WHERE instrument_id = $1 AND valid_until > $2
	`, discountsTable)

	err := s.db.Get(&discount, query, instrumentId, time.Now())

	if err == nil {
		return &discount, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	slog.Error("unknown internal server error during discount obtaining instrument",
		"err", err.Error())
	return nil, &httpError.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "internal server error during discount obtaining instrument",
	}
}

func (s *InstPostgres) DeleteInstrument(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE instrument_id = $1", instrumentsTable)
	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("instrument with id %d not found", id),
		}
	}

	return nil
}

func (s *InstPostgres) GetCategories() ([]entity.Category, error) {
	var categories []entity.Category
	query := fmt.Sprintf(`SELECT category_id, category_name, category_description FROM %s`, categoriesTable)

	err := s.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *InstPostgres) GetManufacturers() ([]entity.Manufacturer, error) {
	var manufacturers []entity.Manufacturer
	query := fmt.Sprintf(`SELECT manufacturer_id, manufacturer_name FROM %s`, manufacturersTable)

	err := s.db.Select(&manufacturers, query)
	if err != nil {
		return nil, err
	}

	return manufacturers, nil
}

func (s *InstPostgres) CreateCategory(category entity.Category) (int, error) {
	var id int
	query := fmt.Sprintf(`
		INSERT INTO %s (category_name, category_description) 
		VALUES ($1, $2) RETURNING category_id
	`, categoriesTable)

	err := s.db.QueryRow(query, category.CategoryName, category.CategoryDescription).Scan(&id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			return 0, &httpError.ErrorWithStatusCode{
				HTTPStatus: http.StatusBadRequest,
				Msg:        "категория с таким названием уже существует",
			}
		}
		return 0, err
	}

	return id, nil
}

func (s *InstPostgres) CreateManufacturer(manufacturer entity.Manufacturer) (int, error) {
	var id int
	query := fmt.Sprintf(`
		INSERT INTO %s (manufacturer_name) 
		VALUES ($1) RETURNING manufacturer_id
	`, manufacturersTable)

	err := s.db.QueryRow(query, manufacturer.ManufacturerName).Scan(&id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
			return 0, &httpError.ErrorWithStatusCode{
				HTTPStatus: http.StatusBadRequest,
				Msg:        "производитель с таким названием уже существует",
			}
		}
		return 0, err
	}

	return id, nil
}
