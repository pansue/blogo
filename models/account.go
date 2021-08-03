package models

import (
	"blogo/global"
	"errors"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string `gorm:"type:VARCHAR(50) NOT NULL"`
	Password string `gorm:"type:VARCHAR(255) NOT NULL"`
}

// CheckAuth verify an account
func CheckAuth(username, password string) (bool, error) {
	var account Account
	err := global.GORM.Select("id").Where(Account{Username: username, Password: password}).First(&account).Error
	if err != nil && nil != errors.New("recode not found") {
		return false, err
	}
	if account.ID > 0 {
		return true, nil
	}
	return false, nil
}

// AddAccount add a single account
func AddAccount(username, password string) error {
	account := Account{
		Username: username,
		Password: password,
	}
	if err := global.GORM.Create(&account).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAccount delete a single account
func DeleteAccount(id int) error {
	if err := global.GORM.Where("id = ?", id).Delete(&Account{}).Error; err != nil {
		return err
	}
	return nil
}

// EditAccount  更新多个属性，只会更新那些被修改了的和非空的字段
func EditAccount(id int, data interface{}) error {
	if err := global.GORM.Model(&Account{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// GetAccount get an account
func GetAccount(id int) (*Account, error) {
	var account Account
	err := global.GORM.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
