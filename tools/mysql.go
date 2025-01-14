package tools

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const mysqlConnectIp = "ihome.sxxpqp.top:3306"

type Product struct {
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Code  string
	Price uint
}
type User struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"unique;Index"`
	PassWord string
}

func Db() *gorm.DB {

	db, err := gorm.Open(mysql.Open("root:Xl123456..@tcp("+mysqlConnectIp+")/test"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&Product{})
	// 自动迁移
	db.AutoMigrate(&User{})
	return db
}
