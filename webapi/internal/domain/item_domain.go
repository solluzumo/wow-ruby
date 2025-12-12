package domain

type ItemDomain struct {
	ID            string
	Name          string
	Price         float32
	RequiredLevel int
	MaxStack      int
	Rarity        string
	ItemType      string
	QuestId       int
	QuestName     string
	Weapon        *WeaponDomain
	Armor         *ArmorDomain
}
