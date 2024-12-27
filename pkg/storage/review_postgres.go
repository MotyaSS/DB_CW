package storage

import (
	"fmt"
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/jmoiron/sqlx"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func newReviewPostgres(db *sqlx.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}

func (s *ReviewPostgres) GetAllReviews(instrumentId int) ([]entity.Review, error) {
	reviews := make([]entity.Review, 0)
	query := fmt.Sprintf(`
		SELECT r.* FROM %s r
		JOIN %s ren ON r.rental_id = ren.rental_id
		WHERE ren.instrument_id = $1
		ORDER BY r.review_id DESC`,
		reviewsTable, rentalsTable)

	err := s.db.Select(&reviews, query, instrumentId)
	return reviews, err
}

func (s *ReviewPostgres) GetReview(id int) (entity.Review, error) {
	var review entity.Review
	query := fmt.Sprintf(`SELECT * FROM %s WHERE review_id = $1`, reviewsTable)
	err := s.db.Get(&review, query, id)
	return review, err
}

func (s *ReviewPostgres) CreateReview(callerId int, review entity.Review) (int, error) {
	// Check if the rental belongs to the caller
	var count int
	err := s.db.Get(&count, fmt.Sprintf(`
		SELECT COUNT(*) FROM %s 
		WHERE rental_id = $1 AND user_id = $2`,
		rentalsTable), review.RentalId, callerId)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "you can only review your own rentals",
		}
	}

	// Check if a review for this rental already exists
	err = s.db.Get(&count, fmt.Sprintf(`
		SELECT COUNT(*) FROM %s 
		WHERE rental_id = $1`,
		reviewsTable), review.RentalId)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusBadRequest,
			Msg:        "review for this rental already exists",
		}
	}

	var id int
	query := fmt.Sprintf(`
		INSERT INTO %s (rental_id, review_text, rating)
		VALUES ($1, $2, $3)
		RETURNING review_id`,
		reviewsTable)

	row := s.db.QueryRow(query, review.RentalId, review.ReviewText, review.Rating)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ReviewPostgres) DeleteReview(callerId int, reviewId int) error {
	result, err := s.db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE review_id = $1`, reviewsTable), reviewId)
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
			Msg:        "review not found",
		}
	}

	return nil
}
