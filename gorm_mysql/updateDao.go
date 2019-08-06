package gorm_mysql

import (
	. "awesomePet/models"
	"fmt"
)

func UpdateUserPassword(user *User) error {
	if err := db.Model(&User{}).Updates(*user).Error; err != nil {
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

func UpdateUserInfo(uid uint64, userInfo *UserInfo) error {
	fmt.Println(uid)
	if err := db.Model(userInfo).Omit("uid").Updates(userInfo).Error; err != nil {
		return err
	}
	return nil
}
