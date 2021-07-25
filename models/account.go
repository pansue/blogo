package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `gorm:"type:VARCHAR(50) NOT NULL"`
	Password string `gorm:"type:VARCHAR(255) NOT NULL"`
}
