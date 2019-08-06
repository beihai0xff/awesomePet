package gorm_mysql

import (
	. "awesomePet/models"
	"fmt"
)

func HasUser(uid uint64) bool {
	return !db.Where("uid = ?", uid).First(&User{}).RecordNotFound()
}

func CreateAccount(user *User, userInfo *UserInfo) error {
	tx := db.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(userInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func CreateUser(user *User) error {
	fmt.Println(*user)
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserInfo(userInfo *UserInfo) error {
	if err := db.Create(userInfo).Error; err != nil {
		return err
	}
	return nil
}
