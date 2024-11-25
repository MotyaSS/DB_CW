package service

import (
	"crypto/sha256"
	"fmt"
	"time"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
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

func (s *AuthService) CreateUser(callerJWT string, user entity.User) (int, error) {
	//Get role id here and check if user can create
	user.Password = generatePasswordHash(user.Password)
	role, err := s.storage.GetRole(user.RoleId)
	if err != nil {
		return 0, fmt.Errorf("role doesn't exist")
	}

	callerId := s.ParseToken(callerJWT)
	switch role.RoleName {
	case "customer":
		return s.CreateCustomer(user)
	case "staff":
		return s.CreateStaff(callerId, user)
	case "chief":
		return s.CreateChief(callerId, user)
	case "admin":
		return s.CreateAdmin(callerId, user)
	default:
		return 0, fmt.Errorf("no handler for such role")
	}
}

//
// TODO: WHATS BELOW
//

func (s *AuthService) CreateCustomer(user entity.User) (int, error) {
	//Get role id here and check if user can create
	user.RoleId = 1
	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func (s *AuthService) CreateStaff(callerId int, user entity.User) (int, error) {
	//Get role id here and check if user can create
	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func (s *AuthService) CreateChief(callerId int, user entity.User) (int, error) {
	//Get role id here and check if user can create
	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func (s *AuthService) CreateAdmin(callerId int, user entity.User) (int, error) {
	//Get role id here and check if user can create
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

func (s *AuthService) ParseToken(token string) int {
	return 0
}
