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
	"github.com/lib/pq"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres) CreateUser(user entity.User) (int, error) {
	slog.Debug("CreateUser invoked", "user", []string{user.Username, user.Email})
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, phone_number, password_hash, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING user_id", usersTable)
	row := s.db.QueryRow(query, user.Username, user.Email, user.PhoneNumber, user.Password, user.RoleId)
	err := row.Scan(&id)
	if err == nil {
		return id, nil
	}
	var pqErr *pq.Error
	ok := errors.As(err, &pqErr)
	if !ok {
		return 0, err
	}

	switch pqErr.Code.Name() {
	case "unique_violation":
		return id, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusBadRequest,
			Msg:        "user with this username already exists",
		}
	}
	return id, err
}

func (s *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	slog.Debug("GetUser invoked", "username", username)
	var user entity.User
	err := s.db.Get(
		&user,
		fmt.Sprintf("SELECT user_id FROM %s WHERE username=$1 AND password_hash=$2", usersTable),
		username,
		password,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return user, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusBadRequest,
			Msg:        "incorrect username or password",
		}
	}

	return user, err
}

func (s *AuthPostgres) GetRole(roleId int) (entity.Role, error) {
	var role entity.Role
	err := s.db.Get(
		&role,
		fmt.Sprintf("SELECT * FROM %s WHERE role_id=$1", rolesTable),
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

func (s *AuthPostgres) GetUserRole(userId int) (entity.Role, error) {
	var roleId int
	// Сначала получаем role_id пользователя
	err := s.db.Get(
		&roleId,
		fmt.Sprintf("SELECT role_id FROM %s WHERE user_id = $1", usersTable),
		userId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Role{}, &httpError.ErrorWithStatusCode{
				HTTPStatus: http.StatusNotFound,
				Msg:        "user not found",
			}
		}
		return entity.Role{}, err
	}

	// Затем получаем полную информацию о роли
	var role entity.Role
	err = s.db.Get(
		&role,
		fmt.Sprintf("SELECT role_id, role_name FROM %s WHERE role_id = $1", rolesTable),
		roleId,
	)
	if err != nil {
		return entity.Role{}, err
	}

	return role, nil
}

func (s *AuthPostgres) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	query := fmt.Sprintf(`
		SELECT user_id, username, email, phone_number, role_id 
		FROM %s
	`, usersTable)

	err := s.db.Select(&users, query)
	if err != nil {
		return nil, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "error getting users list",
		}
	}

	return users, nil
}

func (s *AuthPostgres) GetUserById(userId int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf(`
		SELECT u.user_id, u.username, u.email, u.phone_number, u.role_id 
		FROM %s u WHERE u.user_id = $1
	`, usersTable)

	err := s.db.Get(&user, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, &httpError.ErrorWithStatusCode{
				HTTPStatus: http.StatusNotFound,
				Msg:        "user not found",
			}
		}
		return user, err
	}

	return user, nil
}

func (s *AuthPostgres) DeleteUser(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersTable)
	result, err := s.db.Exec(query, userId)
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
			Msg:        "user not found",
		}
	}

	return nil
}
