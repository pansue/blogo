package main

import (
	"github.com/Dark15/blogo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:rootroot@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err := db.AutoMigrate(&models.Account{}, &models.Article{}, &models.Tag{}, &models.Friend{})

	if err != nil {
		return
	}
}
