package models

import (
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//蛇形小写
type User struct {
	Model
	Uid  uint64 `gorm:"primary_key"`
	Salt string `gorm:"type:char(64); not null"`
	Key  string `gorm:"type:char(64); not null"`
}

type UserInfo struct {
	Model
	Uid         uint64 `json:"uid"  gorm:"primary_key"`
	Nickname    string `json:"nickname" gorm:"type:varchar(25); not null"`
	Sex         uint64 `json:"sex" gorm:"column:sex; not null"`
	Description string `json:"description" gorm:"type:varchar(25)"`
	Email       string `json:"email" gorm:"type:varchar(25); not null"`
	City        uint64 `json:"city"`
	Street      string `json:"street" gorm:"type:varchar(64)"`
	Ext         string `json:"ext" gorm:"type:varchar(6); not null"`
}
