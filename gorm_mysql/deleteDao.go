package gorm_mysql

import (
	. "awesomePet/models"
)

func DeleteUser(uid uint64) (string, error) {
	tx := db.Begin()
	var ext string
	if err := db.Table("user_info").Select("ext").Where("uid = ?", uid).Scan(&ext).Error; err != nil {
		return ext, err
	}
	if err := tx.Delete(&UserInfo{Uid: uid}).Error; err != nil {
		tx.Rollback()
		return ext, err
	}
	if err := tx.Delete(&User{Uid: uid}).Error; err != nil {
		tx.Rollback()
		return ext, err
	}
	return ext, tx.Commit().Error
}

func DeleteBlog(pet *Pet) (err error) {
	tx := db.Begin()
	if err := tx.Delete(pet).Error; err != nil {
		tx.Rollback()
		return
	}
	if err := tx.Where("refer_id = ?", pet.ID).Delete(&Pic{}).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}
