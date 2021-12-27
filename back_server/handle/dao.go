package handle

import (
	"back_server/db"
	"back_server/model"
)

func ValidUser(u *model.User) error {
	u.Password = EncodePassword(u.Password)
	return db.DB.Model(&model.User{}).Where("username = ? and password = ?", u.Username, u.Password).First(&model.User{}).Error
}

func UpdateToken(u *model.User, token string) error {
	return db.DB.Model(&model.User{}).Where(u).Update("token", token).Error
}

func ExistUser(u *model.User) error {
	return db.DB.Model(&model.User{}).Where("username = ?", u.Username).First(&model.User{}).Error
}

func AddUser(u *model.User) error {
	u.Password = EncodePassword(u.Password)
	return db.DB.Create(u).Error
}
