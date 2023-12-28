package web

type UserUpdateRequest struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Username string            `json:"username"`
	Roles    []UserRoleRequest `json:"roles" validate:"required"`
}
