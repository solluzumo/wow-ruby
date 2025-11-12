package models

type User struct {
	*BaseModel
	Login string `db:"login"`
	Hash  string `db:"hash"`
}
