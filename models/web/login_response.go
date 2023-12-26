package web

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// NewLoginResponse returns new LoginResponse.
func NewLoginResponse(AccessToken string) *LoginResponse {

	return &LoginResponse{
		AccessToken: AccessToken,
	}
}
