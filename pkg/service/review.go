package service

import (
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/MotyaSS/DB_CW/pkg/storage"
)

type ReviewService struct {
	storage storage.Review
	auth    Authorisation
}

func NewReviewService(storage storage.Review, auth Authorisation) *ReviewService {
	return &ReviewService{
		storage: storage,
		auth:    auth,
	}
}

func (s *ReviewService) GetAllReviews(instrumentId int) ([]entity.Review, error) {
	return s.storage.GetAllReviews(instrumentId)
}

func (s *ReviewService) GetReview(id int) (entity.Review, error) {
	return s.storage.GetReview(id)
}

func (s *ReviewService) CreateReview(callerId int, review entity.Review) (int, error) {
	// Check if user has customer or higher permissions
	userRole, err := s.auth.GetUserRole(callerId)
	if err != nil {
		return 0, err
	}

	if userRole.RoleId < entity.RoleCustomer.RoleId {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "insufficient permissions to create reviews",
		}
	}

	return s.storage.CreateReview(callerId, review)
}

func (s *ReviewService) DeleteReview(callerId int, reviewId int) error {
	// Check if user has staff or higher permissions
	userRole, err := s.auth.GetUserRole(callerId)
	if err != nil {
		return err
	}

	if userRole.RoleId < entity.RoleStaff.RoleId {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "insufficient permissions to delete reviews",
		}
	}

	return s.storage.DeleteReview(callerId, reviewId)
}
