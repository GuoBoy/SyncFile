package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string `json:"username,omitempty" gorm:"column:username" form:"username" binding:"required"`
	Password string `json:"password,omitempty" gorm:"column:password" form:"password" binding:"required"`
	Token string `gorm:"column:token"`
}
