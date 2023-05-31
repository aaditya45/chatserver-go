package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string    `json : "username" gorm:"unique" gorm:"primaryKey" binding : "min=2,max=12"`
	Password    string       `json : "password" binding : "min=2,max=12"`
}
