package models

import (
	"github.com/solluzumo/wow-ruby/pkg/enums"
)

type Item struct {
	*BaseModel
	Name          string             `db:"name"`
	Price         float32            `db:"price"`
	RequiredLevel int                `db:"required_level"`
	MaxStack      int                `db:"max_stack"`
	Rarity        enums.RarityEnum   `db:"rarity"`
	ItemType      enums.ItemTypeEnum `db:"item_type"`
	QuestId       int                `db:"quest_id"`
	QuestName     string             `db:"quest_name"`
}
