package dto

type UserCreateRequest struct {
	Login    string
	Password string
}

type CreateUserDto struct {
	Login string
	Hash  string
}
