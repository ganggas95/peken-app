package web

import "peken-be/models/domain"

type UserRoleResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserResponse struct {
	Id       uint               `json:"id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Roles    []UserRoleResponse `json:"roles"`
}

func NewUserResponse(user domain.User) UserResponse {
	var roles []UserRoleResponse
	for _, role := range user.Roles {
		roles = append(roles, UserRoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}
	return UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
		Roles:    roles,
	}
}
