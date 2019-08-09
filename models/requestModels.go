package models

type RequestUser struct {
	Uid         uint64 `json:"uid"`
	UserName    string `json:"username"`
	Sex         uint8  `json:"sex"`
	Description string `json:"description"`
	Email       string `json:"email"`
	City        uint64 `json:"city"`
	Street      string `json:"street"`
	Password    string `json:"password"`
}

type PasswordReset struct {
	Uid         uint64 `json:"uid"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

type RequestPet struct {
	ID          uint   `json:"id" gorm:"column:id"`
	Title       string `json:"title" gorm:"type:varchar(15)"`
	Description string `json:"description" gorm:"type:varchar(140)"`
	Tag         string `json:"tag" gorm:"type:varchar(32)"` // json
}
