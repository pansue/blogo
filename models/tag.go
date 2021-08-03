package models

import (
	"blogo/global"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(255) NOT NULL"`
}

// AddTag Add a Tag
func AddTag(name string) error {
	tag := Tag{
		Name: name,
	}
	if err := global.GORM.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

// DeleteTag delete a tag
func DeleteTag(id int) error {
	if err := global.GORM.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := global.GORM.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := global.GORM.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetTags gets a list of tags
func GetTags() ([]Tag, error) {
	var tags []Tag
	err := global.GORM.Select("*").Find(&tags).Error
	if err != nil {
		return tags, err
	} else {
		return nil, err
	}
}
