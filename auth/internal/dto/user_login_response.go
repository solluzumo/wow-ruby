package dto

type UserLoginResponse struct {
	AcessToken   string `json:"acess_token"`
	RefreshToken string `json:"refresh_token"`
	IsActive     bool   `json:"is_active"`
}
