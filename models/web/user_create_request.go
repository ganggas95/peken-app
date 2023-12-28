package web

type UserRoleRequest struct {
	RoleID uint `json:"role_id" validate:"required"`
}

type UserCreateRequest struct {
	Name     string            `json:"name" validate:"required"`
	Email    string            `json:"email" validate:"required,email"`
	Password string            `json:"password" validate:"required,min=8,max=32"`
	Username string            `json:"username" validate:"required,min=3,max=32"`
	Roles    []UserRoleRequest `json:"roles" validate:"required"`
}
