package config

import (
	"apigo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1)/log_barang?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	MigrateDB()
}

func MigrateDB() {
	DB.AutoMigrate(&model.Admin{})
}

// func ReadDB() {
// 	DB.Create(&model.Admin{Name: "admin", })
// }