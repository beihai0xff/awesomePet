package gorm_mysql

import (
	. "awesomePet/models"
)

func UpdateUserPassword(user *User) error {
	return db.Model(&User{}).Updates(*user).Error
}

func UpdateUserInfo(userInfo *UserInfo) error {
	return db.Model(userInfo).Omit("uid").Updates(userInfo).Error
}

func UpdatePPExt(uid uint64, ext string) error {
	return db.Model(&User{}).Where("uid = ?", uid).Update("ext", ext).Error
}

func UpdatePet(uid uint64, updatePet *Pet) error {
	return db.Model(updatePet).Where("uid = ?", uid).Updates(updatePet).Error
}
