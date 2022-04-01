package dataAccess

import (
	"GoRepo/source/assignment/models"
	"gorm.io/gorm"
)

type Chart_da interface {
	GetAllCharts(userId string) ([]models.Chart, error)
	GetChartById(chartId string) (models.Chart, error)
	GetFavouriteCharts(userId string) ([]models.Chart, error)
	SetFavouriteChart(chartId string) (bool, error)
	UnsetFavouriteChart(chartId string) (bool, error)
	DeleteChart(chartId string) (bool, error)
}

type chart_da struct {
	db gorm.DB
}

func NewChart_da(db gorm.DB) Chart_da {
	return &chart_da{db: db}
}

func (d *chart_da) GetAllCharts(userId string) ([]models.Chart, error) {
	var charts []models.Chart
	d.db.Where("userId = ?", userId).Find(&charts)
	return charts, nil
}

func (d *chart_da) GetChartById(chartId string) (models.Chart, error) {
	var chart models.Chart
	d.db.First(&chart, chartId)
	return chart, nil
}

func (d *chart_da) GetFavouriteCharts(userId string) ([]models.Chart, error) {
	var charts []models.Chart
	d.db.Where("favourite = true AND userId = ?", userId).Find(&charts)
	return charts, nil
}

func (d *chart_da) SetFavouriteChart(chartId string) (bool, error) {
	err := d.db.Model(&models.Chart{}).Where("id = ?", chartId).Update("favourite", true).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *chart_da) UnsetFavouriteChart(chartId string) (bool, error) {
	err := d.db.Model(&models.Chart{}).Where("id = ?", chartId).Update("favourite", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *chart_da) DeleteChart(chartId string) (bool, error) {
	err := d.db.Model(&models.Chart{}).Delete("id = ?", chartId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
