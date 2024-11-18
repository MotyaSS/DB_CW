package DB_CW

type User struct {
	UserId      int    `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	RoleId      string `json:"role"`
}

type Role struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}
