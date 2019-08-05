package models

//蛇形小写
type User struct {
	Uid      uint64 `json:"uid"  gorm:"primary_key"`
	UserName string `json:"username" gorm:"column:username; type:varchar(25); not null"`
	Salt     string `json:"salt" gorm:"type:char(64); not null"`
	Key      string `json:"key"  gorm:"type:char(64); not null"`
}