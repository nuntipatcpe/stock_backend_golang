package model

import "time"

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique" from:"username" binding:"required"`
	Password string `form:"passwoed" binding:"required"`
	Level    string `gorm:"defalult:normal"`
	CreateAt time.Time
}
