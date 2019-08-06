package gorm_mysql

import "awesomePet/models"

func UpdateUserPassword(user models.User) error {
	if err := db.Model(&models.User{}).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePPExt(uid uint64, ext string) error {
	if err := db.Model(&models.User{}).Where("uid = ?", uid).Update("ext", ext).Error; err != nil {
		return err
	}
	return nil
}
