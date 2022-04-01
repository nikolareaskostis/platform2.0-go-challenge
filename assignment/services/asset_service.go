package services

import (
	dataAccess "GoRepo/source/assignment/dataAccess"
	models "GoRepo/source/assignment/models"
)

type AssetService interface {
	GetAllAssets(userId string) ([]models.Asset, error)
	GetFavouriteAssets(userId string) ([]models.Asset, error)
}

type assetService struct {
	asset_da dataAccess.Asset_da
}

func NewAssetService(asset_da dataAccess.Asset_da) AssetService {
	return &assetService{asset_da: asset_da}
}

func (s *assetService) GetAllAssets(userId string) ([]models.Asset, error) {
	assets := make([]models.Asset, 0)
	assets, err := s.asset_da.GetAllAssets(userId)

	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (s *assetService) GetFavouriteAssets(userId string) ([]models.Asset, error) {
	assets := make([]models.Asset, 0)
	assets, err := s.asset_da.GetFavouriteAssets(userId)

	if err != nil {
		return nil, err
	}

	return assets, nil
}