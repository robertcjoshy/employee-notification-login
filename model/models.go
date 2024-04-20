package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (user *User) Save() error {
	err := Database.Model(&user).Create(&user).Error
	return err
}
