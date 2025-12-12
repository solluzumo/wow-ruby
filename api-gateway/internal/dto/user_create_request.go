package dto

type UserCreateRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
