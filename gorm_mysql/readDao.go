package gorm_mysql

import (
	. "awesomePet/models"
)

func GetUserPassword(m *User) error {
	return db.First(m).Error
}

func GetUserInfo(m *UserInfo) error {
	return db.First(m).Error
}

func GetUserBlog(uid *uint64) (*[]Pet, error) {
	var m []Pet
	if err := db.Where("uid = ?", uid).Find(&m).Error; err != nil {
		return &m, err
	}
	for i := range m {
		if err := db.Model(&m[i]).Related(&m[i].Pic, "refer_id").Error; err != nil {
			return &m, err
		}
	}
	return &m, nil
}

func GetBlogById(pet *Pet) error {
	return db.First(&pet).Related(&pet.Pic, "refer_id").Error
}
