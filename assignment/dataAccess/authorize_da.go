package dataAccess

import (
	"GoRepo/source/assignment/models"
	"gorm.io/gorm"
)

type Authorize_da interface {
	Login(user models.User) (bool, error)
	Register(user models.User) (bool, error)
	Check(user models.User) (bool, error)
}

type authorize_da struct {
	db gorm.DB
}

func NewAuthorize_da(db gorm.DB) Authorize_da {
	return &authorize_da{db: db}
}

func (d *authorize_da) Login(user models.User) (bool, error) {
	err := d.db.First(&user).Where("id = ? and password = ?", user.Id, user.Password).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *authorize_da) Register(user models.User) (bool, error) {
	err := d.db.Create(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *authorize_da) Check(user models.User) (bool, error) {
	err := d.db.First(&user).Where("id = ? and password = ?", user.Id, user.Password).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
