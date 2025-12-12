package models

type BaseModel struct {
	ID string `db:"id"`
}

func NewBaseModel(id string) *BaseModel {
	return &BaseModel{
		ID: id,
	}
}
