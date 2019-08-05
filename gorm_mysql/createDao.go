package gorm_mysql

func Has(uid uint64) bool {
	return !db.Where("uid = ?", uid).First(&User{}).RecordNotFound()
}

func CreateUser(user *User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
