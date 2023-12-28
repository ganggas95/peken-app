package web

import "peken-be/models/domain"

type UserRoleResponse struct {
	RoleID   uint   `json:"id"`
	RoleName string `json:"name"`
}

type UserResponse struct {
	Id       uint               `json:"id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Name     string             `json:"name"`
	Roles    []UserRoleResponse `json:"roles"`
}

func NewUserResponse(user domain.User) UserResponse {
	var roles []UserRoleResponse
	for _, role := range user.Roles {
		roles = append(roles, UserRoleResponse{
			RoleID:   role.ID,
			RoleName: role.Name,
		})
	}
	return UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
		Name:     user.Name,
		Roles:    roles,
	}
}
