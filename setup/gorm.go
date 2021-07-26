package setup

import (
	"blogo/global"
	"blogo/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", global.CONF.Mysql.User, global.CONF.Mysql.Password, global.CONF.Mysql.Host, global.CONF.Mysql.Database, global.CONF.Mysql.Charset)

	if db, err := gorm.Open(mysql.Open(dsn)); err != nil {
		fmt.Println("mysql error: ", err)
		return nil
	} else {
		setupMysql(db)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(global.CONF.Mysql.MaxIdleConnections)
		sqlDB.SetMaxOpenConns(global.CONF.Mysql.MaxOpenConnections)
		return db
	}
}

func setupMysql(db *gorm.DB) {
	err := db.AutoMigrate(models.Account{}, models.Tag{}, models.Article{}, models.Friend{})
	if err != nil {
		fmt.Println("mysql tables register failed: ", err)
		os.Exit(0)
	}
}
