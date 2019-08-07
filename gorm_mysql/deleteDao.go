package gorm_mysql

import (
	"awesomePet/models"
)

func DeleteUser(uid uint64) (string, error) {
	tx := db.Begin()
	var ext string
	if err := db.Table("user_info").Select("ext").Where("uid = ?", uid).Scan(&ext).Error; err != nil {
		return ext, err
	}
	if err := tx.Delete(&models.UserInfo{Uid: uid}).Error; err != nil {
		tx.Rollback()
		return ext, err
	}
	if err := tx.Delete(&models.User{Uid: uid}).Error; err != nil {
		tx.Rollback()
		return ext, err
	}
	return ext, tx.Commit().Error
}
