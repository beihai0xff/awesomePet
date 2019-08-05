package gorm_mysql

import (
	"awesomePet/api/debug"
	. "awesomePet/models"
)

func GetUserPassword(uid *uint64) *User {
	m := new(User)
	err := db.Where("uid = ?", uid).First(m).Error
	debug.PrintErr(err)
	return m
}
