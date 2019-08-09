package gorm_mysql

import (
	. "awesomePet/models"
)

func UpdateUserPassword(user *User) error {
	if err := db.Model(&User{}).Updates(*user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserInfo(userInfo *UserInfo) error {
	if err := db.Model(userInfo).Omit("uid").Updates(userInfo).Error; err != nil { // 忽略uid
		return err
	}
	return nil
}

func UpdatePPExt(uid uint64, ext string) error {
	if err := db.Model(&User{}).Where("uid = ?", uid).Update("ext", ext).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePet(uid uint64, updatePet *Pet) error {
	if err := db.Model(updatePet).Where("uid = ?", uid).Updates(updatePet).Error; err != nil {
		return err
	}
	return nil
}
