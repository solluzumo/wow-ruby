package models

import "wow-ruby/internal/pkg/enums"

type Weapon struct {
	*BaseModel
	Name          string               `db:"name"`
	Price         float32              `db:"price"`
	RequiredLevel int                  `db:"required_level"`
	MaxStack      int                  `db:"max_stack"`
	Rarity        enums.RarityEnum     `db:"rarity"`
	ItemType      enums.ItemTypeEnum   `db:"item_type"`
	Slot          string               `db:"slot"`
	Durability    int                  `db:"durability"`
	Damage        string               `db:"damage"`
	Speed         float32              `db:"speed"`
	WeaponType    enums.WeaponTypeEnum `db:"weapon_type"`
	SetName       string               `db:"set_name"`
	QuestId       int                  `db:"quest_id"`
	QuestName     string               `db:"quest_name"`
}
