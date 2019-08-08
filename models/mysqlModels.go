package models

import (
	"time"
)

// 蛇形小写
type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"` // 软删除
}

type User struct {
	Uid  uint64 `gorm:"primary_key"`
	Salt string `gorm:"type:char(64); not null"`
	Key  string `gorm:"type:char(64); not null"`
}

type UserInfo struct {
	Model
	Uid         uint64 `json:"uid"  gorm:"primary_key"`
	Nickname    string `json:"nickname" gorm:"type:varchar(25); not null"`
	Sex         uint8  `json:"sex" gorm:"column:sex; not null"`
	Description string `json:"description" gorm:"type:varchar(25)"`
	Email       string `json:"email" gorm:"type:varchar(25); not null"`
	City        uint64 `json:"city"`
	Street      string `json:"street" gorm:"type:varchar(64)"`
	Ext         string `json:"ext" gorm:"type:varchar(6); not null"`
}

type Pet struct {
	Model
	ID          uint   `json:"id" gorm:"column:id"`
	Uid         uint64 `json:"uid" gorm:"index"`
	Title       string `json:"title" gorm:"type:varchar(15)"`
	Description string `json:"description" gorm:"type:varchar(140)"`
	Pic         []Pic  `json:"pic" gorm:"ForeignKey:ReferID"`
	Tag         string `json:"tag" gorm:"type:varchar(32)"` // 传过来json
	Star        uint   `json:"star"`
}

type Pic struct {
	OrderID uint   `json:"orderId"`
	ReferID uint   `json:"referId" gorm:"column:refer_id"`
	PetHash string `json:"petHash" gorm:"type:char(64)"`
	Ext     string `json:"ext" gorm:"type:varchar(6); not null"`
}
