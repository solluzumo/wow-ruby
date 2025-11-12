package models

type Tabler interface {
	TableName() string
}

func (User) TableName() string   { return "users" }
func (Armor) TableName() string  { return "armor" }
func (Weapon) TableName() string { return "weapon" }
func (Item) TableName() string   { return "item" }
func (Quest) TableName() string  { return "quest" }
func (Npc) TableName() string    { return "npc" }
