package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int8   `json:"level"`
	Price       int8   `json:"price"`
	OwnerID     int64  `json:"owner_id"`
}
