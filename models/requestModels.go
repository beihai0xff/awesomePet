package models

type RequestUser struct {
	Uid         uint64 `json:"uid"  gorm:"primary_key"`
	UserName    string `json:"username" gorm:"type:varchar(25); not null"`
	Sex         uint64 `json:"sex" gorm:"column:sex; not null"`
	Description string `json:"description" gorm:"type:varchar(25); not null"`
	Email       string `json:"email" gorm:"type:varchar(25); not null"`
	City        uint64 `json:"city" gorm:"not null"`
	Street      string `json:"street" gorm:"type:varchar(64); not null"`
	Password    string `json:"password"`
}

type PasswordReset struct {
	Uid         uint64 `json:"uid"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
