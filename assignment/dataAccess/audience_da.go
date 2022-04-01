package dataAccess

import (
	"GoRepo/source/assignment/models"
	"gorm.io/gorm"
)

type Audience_da interface {
	GetAllAudiences(userId string) ([]models.Audience, error)
	GetAudienceById(audienceId string) (models.Audience, error)
	GetFavouriteAudiences(userId string) ([]models.Audience, error)
	SetFavouriteAudience(audienceId string) (bool, error)
	UnsetFavouriteAudience(audienceId string) (bool, error)
	DeleteAudience(audienceId string) (bool, error)
}

type audience_da struct {
	db gorm.DB
}

func NewAudience_da(db gorm.DB) Audience_da {
	return &audience_da{db: db}
}

func (d *audience_da) GetAllAudiences(userId string) ([]models.Audience, error) {
	var audiences []models.Audience
	d.db.Where("userId = ?", userId).Find(&audiences)
	return audiences, nil
}

func (d *audience_da) GetAudienceById(audienceId string) (models.Audience, error) {
	var audience models.Audience
	d.db.Where("id = ?", audienceId).Find(&audience)
	return audience, nil
}

func (d *audience_da) GetFavouriteAudiences(userId string) ([]models.Audience, error) {
	var audiences []models.Audience
	d.db.Where("favourite = true AND userId = ?", userId).Find(&audiences)
	return audiences, nil
}

func (d *audience_da) SetFavouriteAudience(audienceId string) (bool, error) {
	err := d.db.Model(&models.Audience{}).Where("id = ?", audienceId).Update("favourite", true).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *audience_da) UnsetFavouriteAudience(audienceId string) (bool, error) {
	err := d.db.Model(&models.Audience{}).Where("id = ?", audienceId).Update("favourite", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *audience_da) DeleteAudience(audienceId string) (bool, error) {
	err := d.db.Model(&models.Audience{}).Delete("id = ?", audienceId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
