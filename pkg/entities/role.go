package entity

type Role struct {
	RoleId   int    `json:"role_id" db:"role_id"`
	RoleName string `json:"role_name" db:"role_name"`
}

var (
	RoleCustomer = Role{RoleId: 1, RoleName: "customer"}
	RoleStaff    = Role{RoleId: 2, RoleName: "staff"}
	RoleChief    = Role{RoleId: 3, RoleName: "chief"}
	RoleAdmin    = Role{RoleId: 4, RoleName: "admin"}
)
