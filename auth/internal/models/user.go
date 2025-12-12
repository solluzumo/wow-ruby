package models

type User struct {
	*BaseModel
	Hash     string `db:"hash"`
	Email    string `db:"email"`
	IsActive bool   `db:"is_active"`
}
