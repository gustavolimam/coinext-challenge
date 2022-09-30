package model

type Inventory struct {
	UserID string `json:"user_id" db:"user_id"`
	Water  int    `json:"water" db:"water"`
	Food   int    `json:"food" db:"food"`
	Drug   int    `json:"drug" db:"drug"`
	Ammo   int    `json:"ammo" db:"ammo"`
}
