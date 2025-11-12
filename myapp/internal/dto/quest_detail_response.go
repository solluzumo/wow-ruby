package dto

type QuestDetailResponse struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	RewardMoney            int    `json:"reward_money"`
	RequiredCharacterLevel int    `json:"required_character_level"`
	QuestLevel             int    `json:"quest_level"`
	Repeatable             bool   `json:"is_repeatable"`
	Difficulty             int    `json:"difficulty"`
}
