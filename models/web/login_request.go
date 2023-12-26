package web

type LoginRequest struct {
	Password string `json:"password" validate:"required,min=8,max=32"`
	Username string `json:"username" validate:"required,min=3,max=32"`
}
