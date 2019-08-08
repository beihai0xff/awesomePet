package gorm_mysql

import (
	"awesomePet/api/debug"
	. "awesomePet/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB // 全局变量用 =

func Init(args *string) {
	var err error
	db, err = gorm.Open("mysql", *args)
	debug.PanicErr(err)
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)
	db.LogMode(true)       // 启用Logger，显示详细日志
	db.SingularTable(true) // 全局禁用表名复数
	fmt.Println("mysql数据库已连接，检查表结构中...")
	CreateTable("User", &User{})
	CreateTable("UserInfo", &UserInfo{})
	CreateTable("Pet", &Pet{})
	CreateTable("Pic", &Pic{})
	//defer db.Close()
}

func CreateTable(name string, table interface{}) {
	if !db.HasTable(table) {
		fmt.Printf("表:%s不存在，正在创建中\n", name)
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(table).Error; err != nil {
			fmt.Printf("表:%s创建失败，请查看原因：\n", name)
			panic(err)
		}
		fmt.Printf("表:%s创建完成\n", name)
	}

}
