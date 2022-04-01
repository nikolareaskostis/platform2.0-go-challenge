package dataAccess

import (
	"GoRepo/source/assignment/models"
	"gorm.io/gorm"
)

type Asset_da interface {
	GetAllAssets(userId string) ([]models.Asset, error)
	GetFavouriteAssets(userId string) ([]models.Asset, error)
}

type asset_da struct {
	db gorm.DB
}

func NewAsset_da(db gorm.DB) Asset_da {
	return &asset_da{db: db}
}

func (d *asset_da) GetAllAssets(userId string) ([]models.Asset, error) {
	var assets []models.Asset
	d.db.Where("userId = ?", userId).Find(&assets)
	return assets, nil
}

func (d *asset_da) GetFavouriteAssets(userId string) ([]models.Asset, error) {
	var assets []models.Asset
	d.db.Where("favourite = true AND userId = ?", userId).Find(&assets)
	return assets, nil
}
