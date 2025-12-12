package dto

type ItemDetailResponse struct {
	ID            string                `json:"id"`
	Name          string                `json:"name"`
	Price         float32               `json:"price"`
	RequiredLevel int                   `json:"required_level"`
	MaxStack      int                   `json:"max_stack"`
	Rarity        string                `json:"rarity"`
	ItemType      string                `json:"item_type"`
	QuestId       int                   `json:"quest_id,omitempty"`
	QuestName     string                `json:"quest_name,omitempty"`
	Weapon        *WeaponDetailResponse `json:"weapon,omitempty"`
	Armor         *ArmorDetailResponse  `json:"armor,omitempty"`
}
