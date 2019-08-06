package gorm_mysql

import (
	. "awesomePet/models"
)

func GetUserPassword(uid *uint64) (*User, error) {
	m := new(User)
	err := db.Where("uid = ?", uid).First(m).Error
	if err != nil {
		return m, err
	}
	return m, nil
}

func GetUserInfo(uid *uint64) (*UserInfo, error) {
	m := new(UserInfo)
	err := db.Where("uid = ?", uid).First(m).Error

	return m, nil
}
