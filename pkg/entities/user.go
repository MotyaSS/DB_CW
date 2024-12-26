package entity

type User struct {
	UserId      int    `json:"user_id" db:"user_id"`
	Username    string `json:"username" db:"username" binding:"required"`
	Email       string `json:"email" db:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" db:"phone_number" binding:"required,e164"`
	Password    string `json:"password" db:"password" binding:"required"`
	RoleId      int    `json:"role_id" db:"role_id"`
}

type Role struct {
	RoleId   int    `json:"role_id" db:"role_id"`
	RoleName string `json:"role_name" db:"role_name"`
}
