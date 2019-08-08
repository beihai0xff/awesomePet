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
	if err := db.Where("uid = ?", uid).First(m).Error; err != nil {
		return m, err
	}
	return m, nil
}

func GetUserBlog(uid *uint64) (*Pet, error) {
	m := new(Pet)
	err := db.Where("uid = ?", uid).Find(&m).Error
	if err != nil {
		return m, err
	}
	err = db.Model(&m).Related(&m.Pic, "refer_id").Error
	if err != nil {
		return m, err
	}
	return m, nil
}
