package model

import "time"

type User struct {
	ID           *string    `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Age          int        `json:"age" db:"age"`
	Gender       string     `json:"gender" db:"gender"`
	LastLocation Location   `json:"location" db:"location"`
	Inventory    *Inventory `json:"inventory,omitempty" db:"inventory,omitempty"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	DeadAt       *time.Time `json:"dead_at,omitempty" db:"dead_at,omitempty"`
}

type Location struct {
	Longitude string `json:"longitude" db:"longitude"`
	Latitude  string `json:"latitude" db:"latitude"`
}
