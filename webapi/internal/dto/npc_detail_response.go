package dto

type NpcDetailResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Health      int    `json:"health"`
	Mana        int    `json:"mana"`
	Level       int    `json:"level"`
	Tameable    bool   `json:"tameable"`
	Faction     string `json:"faction"`
	Reaction    string `json:"reaction"`
	Location    string `json:"location"`
	RespawnTime int    `json:"respawn_time"`
}
