package dataAccess

import (
	"GoRepo/source/assignment/models"
	"gorm.io/gorm"
)

type Insight_da interface {
	GetAllInsights(userId string) ([]models.Insight, error)
	GetInsightById(chartId string) (models.Insight, error)
	GetFavouriteInsights(userId string) ([]models.Insight, error)
	SetFavouriteInsight(chartId string) (bool, error)
	UnsetFavouriteInsight(chartId string) (bool, error)
	DeleteInsight(chartId string) (bool, error)
}

type insight_da struct {
	db gorm.DB
}

func NewInsight_da(db gorm.DB) Insight_da {
	return &insight_da{db: db}
}

func (d *insight_da) GetAllInsights(userId string) ([]models.Insight, error) {
	var insights []models.Insight
	d.db.Where("userId = ?", userId).Find(&insights)
	return insights, nil
}

func (d *insight_da) GetInsightById(insightId string) (models.Insight, error) {
	var insight models.Insight
	d.db.First(&insight, insightId)
	return insight, nil
}

func (d *insight_da) GetFavouriteInsights(userId string) ([]models.Insight, error) {
	var charts []models.Insight
	d.db.Where("favourite = true AND userId = ?", userId).Find(&charts)
	return charts, nil
}

func (d *insight_da) SetFavouriteInsight(insightId string) (bool, error) {
	err := d.db.Model(&models.Insight{}).Where("id = ?", insightId).Update("favourite", true).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *insight_da) UnsetFavouriteInsight(insightId string) (bool, error) {
	err := d.db.Model(&models.Insight{}).Where("id = ?", insightId).Update("favourite", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *insight_da) DeleteInsight(insightId string) (bool, error) {
	err := d.db.Model(&models.Insight{}).Delete("id = ?", insightId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
