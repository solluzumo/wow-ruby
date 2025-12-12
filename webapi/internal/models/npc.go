package models

type Npc struct {
	*BaseModel
	Name        string `db:"name"`
	Health      int    `db:"health"`
	Mana        int    `db:"mana"`
	Level       int    `db:"level"`
	Tameable    bool   `db:"tameable"`
	Faction     string `db:"faction"`
	Reaction    string `db:"reaction"`
	Location    string `db:"location"`
	RespawnTime int    `db:"respawn_time"`
}
