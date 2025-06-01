package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string
	Description string
	Level       uint8
	Price       uint
	OwnerID     uint
}
