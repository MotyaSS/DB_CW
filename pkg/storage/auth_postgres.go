package storage

import (
	"fmt"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, phone_number, password_hash, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING user_id", usersTable)
	row := s.db.QueryRow(query, user.Username, user.Email, user.PhoneNumber, user.Password, user.RoleId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	err := s.db.Get(
		&user,
		fmt.Sprintf("SELECT user_id FROM %s WHERE username=$1 AND password_hash=$2", usersTable),
		username,
		password,
	)
	return user, err
}

func (s *AuthPostgres) GetRole(roleId int) (entity.Role, error) {
	var role entity.Role
	err := s.db.Get(
		&role,
		fmt.Sprintf("SELECT role_name FROM %s WHERE role_id=$1", rolesTable),
		roleId,
	)
	return role, err
}
func (s *AuthPostgres) GetRoleId(roleName string) (int, error) {
	var role int
	err := s.db.Get(
		&role,
		fmt.Sprintf("SELECT role_id FROM %s WHERE role_name=$1", rolesTable),
		roleName,
	)
	return role, err
}

func (s *AuthPostgres) GetAllRoles() ([]entity.Role, error) {
	var result []entity.Role
	err := s.db.Select(
		&result,
		fmt.Sprintf("SELECT role_id, role_name FROM %s", rolesTable),
	)
	return result, err
}
