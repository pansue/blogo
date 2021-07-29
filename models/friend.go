package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	Name      string `gorm:"type:VARCHAR(50) NOT NULL"`
	Link      string `gorm:"type:VARCHAR(255) NOT NULL"`
	AvatarURL string `gorm:"type:VARCHAR(255) NOT NULL"`
	Desc      string `gorm:"type:VARCHAR(255)"`
}
