package models

type Tabler interface {
	TableName() string
}

func (User) TableName() string { return "users" }
