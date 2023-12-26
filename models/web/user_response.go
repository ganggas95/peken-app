package web

import "peken-be/models/domain"

type UserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func NewUserResponse(user domain.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
		Name:     user.Name,
	}
}
