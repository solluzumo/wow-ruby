package models

import (
	"github.com/solluzumo/wow-ruby/pkg/enums"
)

type Armor struct {
	*BaseModel
	Name          string              `db:"name"`
	Price         float32             `db:"price"`
	RequiredLevel int                 `db:"required_level"`
	MaxStack      int                 `db:"max_stack"`
	Rarity        enums.RarityEnum    `db:"rarity"`
	ItemType      enums.ItemTypeEnum  `db:"item_type"`
	Slot          string              `db:"slot"`
	Durability    int                 `db:"durability"`
	ArmorType     enums.ArmorTypeEnum `db:"armor_type"`
	ArmorValue    int                 `db:"armor_value"`
	SetName       string              `db:"set_name"`
	QuestId       int                 `db:"quest_id"`
	QuestName     string              `db:"quest_name"`
}
