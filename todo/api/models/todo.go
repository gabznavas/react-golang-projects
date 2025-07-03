package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
}

func (Todo) TableName() string {
	return "todos"
}
