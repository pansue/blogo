package main

import (
	"blogo/global"
	"blogo/setup"
	"database/sql"
	"fmt"
)

func init() {
	global.VIPER = setup.Viper()
	global.GORM = setup.Gorm()
}

func main() {
	db, _ := global.GORM.DB()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

}
