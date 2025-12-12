package models

type Quest struct {
	*BaseModel
	Name                   string `db:"name"`
	RewardMoney            int    `db:"reward_money"`
	RequiredCharacterLevel int    `db:"required_character_level"`
	QuestLevel             int    `db:"quest_level"`
	Repeatable             bool   `db:"is_repeatable"`
	Difficulty             int    `db:"difficulty"`
}
