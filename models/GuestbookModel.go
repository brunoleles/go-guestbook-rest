// TODO Check anda validates if refactoring to package "database" makes sense
package models

import "gorm.io/gorm"

type GuestbookModel struct {
	ID      uint   `gorm:"column:id;primary key;autoIncrement" json:"ID"`
	Name    string `gorm:"column:name" json:"name"`
	Message string `gorm:"column:message" json:"message"`
	gorm.Model
}
