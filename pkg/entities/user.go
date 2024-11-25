package entity

type User struct {
	UserId      int    `json:"user_id" db:"user_id"`
	Username    string `json:"username" binding:"required" db:"username"`
	Email       string `json:"email" binding:"required,email" db:"email"`
	PhoneNumber string `json:"phone_number" binding:"required" db:"phone_number"`
	Password    string `json:"password" binding:"required" db:"password"`
	RoleId      int    `json:"role_id" db:"role_id"`
}

type Role struct {
	RoleId   int    `json:"role_id" db:"role_id"`
	RoleName string `json:"role_name" db:"role_name"`
}
