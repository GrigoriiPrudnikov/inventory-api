package models

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
	Price       int    `json:"price"`
	OwnerID     int    `json:"owner_id"`
}
