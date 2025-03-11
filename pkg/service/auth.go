package service

import (
	"crypto/sha256"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/MotyaSS/DB_CW/pkg/storage"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "Solevoy!gv13fa788fy0a67sDf4"
	signingKey = "nzo38r9b09&a^1_@)_u09ahj1;a"
	tokenTTL   = time.Hour * 12
)

type AuthService struct {
	storage storage.Authorisation
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(storage storage.Authorisation) *AuthService {
	return &AuthService{storage: storage}
}

// HasPermission checks if the user role has sufficient permissions for the required role
func (s *AuthService) HasPermission(userRole, requiredRole entity.Role) bool {
	// Admin has all permissions
	if userRole.RoleId == entity.RoleAdmin.RoleId {
		return true
	}

	// Chief can manage staff and customers
	if userRole.RoleId == entity.RoleChief.RoleId {
		return requiredRole.RoleId <= entity.RoleChief.RoleId
	}

	// Staff can manage customers
	if userRole.RoleId == entity.RoleStaff.RoleId {
		return requiredRole.RoleId <= entity.RoleStaff.RoleId
	}

	// Customers can only access customer-level permissions
	if userRole.RoleId == entity.RoleCustomer.RoleId {
		return requiredRole.RoleId == entity.RoleCustomer.RoleId
	}

	return false
}

// CheckPermission checks if the user has the required role or higher
func (s *AuthService) CheckPermission(userId int, requiredRole entity.Role) error {
	userRole, err := s.storage.GetRole(userId)
	if err != nil {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error getting user role",
		}
	}

	if !s.HasPermission(userRole, requiredRole) {
		return &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "insufficient permissions",
		}
	}

	return nil
}

func (s *AuthService) GetAllRoles() ([]entity.Role, error) {
	return s.storage.GetAllRoles()
}

func (s *AuthService) GetRole(roleId int) (entity.Role, error) {
	return s.storage.GetRole(roleId)
}

func (s *AuthService) GetUserRole(userId int) (entity.Role, error) {
	role, err := s.storage.GetRole(userId)
	if err != nil {
		return entity.Role{}, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error getting user role",
		}
	}
	return role, nil
}

func (s *AuthService) CreateCustomer(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	if user.RoleId != entity.RoleCustomer.RoleId {
		slog.Info("Role mismatch: expected customer, got different. Created user with customer role", "role", user.RoleId)
	}
	user.RoleId = entity.RoleCustomer.RoleId
	return s.storage.CreateUser(user)
}

func (s *AuthService) CreateUser(callerId int, user entity.User) (int, error) {
	// If callerId is -1, it means we're creating a customer account (self-registration)
	if callerId == -1 {
		if user.RoleId != entity.RoleCustomer.RoleId {
			return 0, &httpError.ErrorWithStatusCode{
				HTTPStatus: http.StatusForbidden,
				Msg:        "self-registration is only allowed for customer accounts",
			}
		}
		user.Password = generatePasswordHash(user.Password)
		return s.storage.CreateUser(user)
	}

	// Check if caller has permission to create users with this role
	callerRole, err := s.GetUserRole(callerId)
	if err != nil {
		return 0, err
	}

	// Only admin can create admin accounts
	if user.RoleId == entity.RoleAdmin.RoleId && callerRole.RoleId != entity.RoleAdmin.RoleId {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "only admins can create admin accounts",
		}
	}

	// Chief and higher can create staff accounts
	if user.RoleId == entity.RoleStaff.RoleId && callerRole.RoleId < entity.RoleChief.RoleId {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "insufficient permissions to create staff accounts",
		}
	}

	// Only admin can create chief accounts
	if user.RoleId == entity.RoleChief.RoleId && callerRole.RoleId != entity.RoleAdmin.RoleId {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusForbidden,
			Msg:        "only admins can create chief accounts",
		}
	}

	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.storage.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.UserId,
		})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&tokenClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(signingKey), nil
		})
	if err != nil {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusUnauthorized,
			Msg:        err.Error(),
		}
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusUnauthorized,
			Msg:        "token claims are not of type *tokenClaims",
		}
	}

	return claims.UserId, nil
}

func (s *AuthService) GetAllUsers() ([]entity.User, error) {
	users, err := s.storage.GetAllUsers()
	if err != nil {
		return nil, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error getting users list",
		}
	}

	// Не возвращаем чувствительные данные
	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}

func (s *AuthService) GetUserById(userId int) (entity.User, error) {
	return s.storage.GetUserById(userId)
}

func (s *AuthService) DeleteUser(userId int) error {
	return s.storage.DeleteUser(userId)
}
