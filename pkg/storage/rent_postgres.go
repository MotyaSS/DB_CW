package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type RentPostgres struct {
	db *sqlx.DB
}

func newRentPostgres(db *sqlx.DB) *RentPostgres {
	return &RentPostgres{db: db}
}

func (r *RentPostgres) CreateRental(rental entity.Rental) (int, error) {
	var id int
	query := fmt.Sprintf(`
		INSERT INTO %s (user_id, instrument_id, rental_date, return_date)
		VALUES ($1, $2, $3, $4)
		RETURNING rental_id
	`, rentalsTable)

	err := r.db.QueryRow(query,
		rental.UserId,
		rental.InstrumentId,
		rental.RentalDate,
		rental.ReturnDate,
	).Scan(&id)

	if err != nil {
		slog.Error("error creating rental", "err", err.Error())
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error creating rental",
		}
	}

	return id, nil
}

func (r *RentPostgres) GetRental(id int) (entity.Rental, error) {
	var rental entity.Rental
	query := fmt.Sprintf(`
		SELECT * FROM %s WHERE rental_id = $1
	`, rentalsTable)

	err := r.db.Get(&rental, query, id)
	if err == nil {
		return rental, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return rental, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("rental with id %d not found", id),
		}
	}

	slog.Error("error getting rental", "err", err.Error())
	return rental, &httpError.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "error getting rental",
	}
}

func (r *RentPostgres) GetUserRentals(userId int) ([]entity.Rental, error) {
	var rentals []entity.Rental
	query := fmt.Sprintf(`
		SELECT r.*, i.instrument_name, i.description, i.price_per_day, i.image_url
		FROM %s r
		LEFT JOIN %s i ON r.instrument_id = i.instrument_id
		WHERE r.user_id = $1
	`, rentalsTable, instrumentsTable)

	type rentalWithInstrument struct {
		entity.Rental
		InstrumentName string          `db:"instrument_name"`
		Description    string          `db:"description"`
		PricePerDay    decimal.Decimal `db:"price_per_day"`
		ImageUrl       string          `db:"image_url"`
	}

	var rentalRows []rentalWithInstrument
	err := r.db.Select(&rentalRows, query, userId)
	if err != nil {
		slog.Error("error getting user rentals", "err", err.Error())
		return nil, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error getting user rentals",
		}
	}

	for _, row := range rentalRows {
		rental := row.Rental
		rental.Instrument = &entity.Instrument{
			InstrumentName: row.InstrumentName,
			Description:    row.Description,
			PricePerDay:    row.PricePerDay,
			ImageUrl:       &row.ImageUrl,
		}
		rentals = append(rentals, rental)
	}

	return rentals, nil
}

func (r *RentPostgres) GetInstrumentRentals(instrumentId int) ([]entity.Rental, error) {
	var rentals []entity.Rental
	query := fmt.Sprintf(`
		SELECT * FROM %s WHERE instrument_id = $1
	`, rentalsTable)

	err := r.db.Select(&rentals, query, instrumentId)
	if err != nil {
		slog.Error("error getting instrument rentals", "err", err.Error())
		return nil, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error getting instrument rentals",
		}
	}

	return rentals, nil
}

func (r *RentPostgres) UpdateRental(rental entity.Rental) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET user_id = $1, instrument_id = $2, rental_date = $3, return_date = $4
		WHERE rental_id = $5
	`, rentalsTable)

	result, err := r.db.Exec(query,
		rental.UserId,
		rental.InstrumentId,
		rental.RentalDate,
		rental.ReturnDate,
		rental.RentalId,
	)

	if err != nil {
		slog.Error("error updating rental", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error updating rental",
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("error getting rows affected", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error updating rental",
		}
	}

	if rowsAffected == 0 {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("rental with id %d not found", rental.RentalId),
		}
	}

	return nil
}

func (r *RentPostgres) DeleteRental(id int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s WHERE rental_id = $1
	`, rentalsTable)

	result, err := r.db.Exec(query, id)
	if err != nil {
		slog.Error("error deleting rental", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error deleting rental",
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("error getting rows affected", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error deleting rental",
		}
	}

	if rowsAffected == 0 {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("rental with id %d not found", id),
		}
	}

	return nil
}

func (r *RentPostgres) ReturnInstrument(rentalId int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET return_date = $1
		WHERE rental_id = $2 AND return_date IS NULL
	`, rentalsTable)

	result, err := r.db.Exec(query, time.Now(), rentalId)
	if err != nil {
		slog.Error("error returning instrument", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error returning instrument",
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("error getting rows affected", "err", err.Error())
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error returning instrument",
		}
	}

	if rowsAffected == 0 {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("active rental with id %d not found", rentalId),
		}
	}

	return nil
}
