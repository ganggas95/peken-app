package web

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewLoginResponse returns new LoginResponse.
func NewLoginResponse(AccessToken string, RefreshToken string) *LoginResponse {

	return &LoginResponse{
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
	}
}
