package gorm_mysql

import (
	. "awesomePet/models"
)

func GetUserPassword(uid *uint64) (*User, error) {
	m := new(User)
	err := db.Where("uid = ?", uid).First(m).Error
	return m, err
}

func GetUserInfo(uid *uint64) (*UserInfo, error) {
	m := new(UserInfo)
	err := db.Where("uid = ?", uid).First(m).Error
	return m, err
}

func GetUserBlog(uid *uint64) (*[]Pet, error) {
	var m []Pet
	err := db.Where("uid = ?", uid).Find(&m).Error
	if err != nil {
		return &m, err
	}
	for i := range m {
		err = db.Model(&m[i]).Related(&m[i].Pic, "refer_id").Error
		if err != nil {
			return &m, err
		}
	}
	return &m, err
}

func GetBlogById(pet *Pet) error {
	return db.First(&pet).Related(&pet.Pic, "refer_id").Error
}
