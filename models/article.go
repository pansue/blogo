package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title          string `gorm:"type:VARCHAR(255) NOT NULL"`
	Desc           string `gorm:"type:VARCHAR(255)"`
	ContentPath    string `gorm:"type:VARCHAR(255) NOT NULL"`
	State          uint   `gorm:"type:TINYINT UNSIGNED NOT NULL;default:0"`
	CoverImageURL  string `gorm:"type:VARCHAR(255) NOT NULL"`
	Views          uint   `gorm:"type:INT UNSIGNED NOT NULL;default:0"`
	EffectiveViews uint   `gorm:"type:INT UNSIGNED NOT NULL;default:0"`
	Tags           []Tag  `gorm:"many2many:article_tags;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}
