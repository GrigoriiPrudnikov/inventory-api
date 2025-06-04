package models

type Item struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       uint8  `json:"level"`
	Price       uint   `json:"price"`
	OwnerID     uint   `json:"owner_id"`
}
