package models

import "gorm.io/gorm"

type GuestbookModel struct {
	ID      uint   `gorm:"primary key;autoIncrement" json:"ID"`
	Name    string `json:"name"`
	Message string `json:"message"`
	gorm.Model
}
