package gorm_mysql

import (
	. "awesomePet/models"
)

func HasUser(uid uint64) bool {
	return !db.Where("uid = ?", uid).First(&User{}).RecordNotFound()
}

func CreateAccount(user *User, userInfo *UserInfo) (err error) {
	tx := db.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return
	}
	if err := tx.Create(userInfo).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

func CreatePet(pet *Pet) (err error) {
	return db.Create(pet).Error
}

func CreateUser(user *User) (err error) {
	return db.Create(user).Error
}

func CreateUserInfo(userInfo *UserInfo) (err error) {
	return db.Create(userInfo).Error
}
