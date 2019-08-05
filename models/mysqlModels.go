package models

//蛇形小写
type User struct {
	Uid  uint64 `gorm:"primary_key"`
	Salt string `gorm:"type:char(64); not null"`
	Key  string `gorm:"type:char(64); not null"`
}

type UserInfo struct {
	Uid         uint64 `json:"uid"  gorm:"primary_key"`
	UserName    string `json:"username" gorm:"type:varchar(25); not null"`
	Sex         uint64 `json:"sex" gorm:"column:sex; not null"`
	Description string `json:"description" gorm:"type:varchar(25); not null"`
	Email       string `json:"email" gorm:"type:varchar(25); not null"`
	City        uint64 `json:"city" gorm:"not null"`
	Street      string `json:"street" gorm:"type:varchar(64); not null"`
}
