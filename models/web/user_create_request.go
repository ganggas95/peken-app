package web

type UserRoleRequest struct {
	ID uint `json:"id" validate:"required"`
}

type UserCreateRequest struct {
	Email    string            `json:"email" validate:"required,email"`
	Password string            `json:"password" validate:"required,min=8,max=32"`
	Username string            `json:"username" validate:"required,min=3,max=32"`
	Roles    []UserRoleRequest `json:"roles" validate:"required"`
}
