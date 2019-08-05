package gorm_mysql

import "awesomePet/models"

func UpdateUserPassword(user models.User) error {
	if err := db.Model(&models.User{}).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

