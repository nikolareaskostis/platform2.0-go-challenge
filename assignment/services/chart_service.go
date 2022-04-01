package services

import (
	"GoRepo/source/assignment/dataAccess"
	"GoRepo/source/assignment/models"
)

type ChartService interface {
	GetAllCharts(userId string) ([]models.Chart, error)
	GetFavouriteCharts(userId string) ([]models.Chart, error)
	GetChartById(chartId string) (models.Chart, error)
	SetFavouriteChart(chartId string) (bool, error)
	UnsetFavouriteChart(chartId string) (bool, error)
	DeleteChart(chartId string) (bool, error)
}

type chartService struct {
	chart_da dataAccess.Chart_da
}

func NewChartService(chart_da dataAccess.Chart_da) ChartService {
	return &chartService{chart_da: chart_da}
}

func (s *chartService) GetAllCharts(userId string) ([]models.Chart, error) {
	charts := make([]models.Chart, 0)
	charts, err := s.chart_da.GetAllCharts(userId)

	if err != nil {
		return nil, err
	}

	return charts, nil
}

func (s *chartService) GetFavouriteCharts(userId string) ([]models.Chart, error) {
	charts := make([]models.Chart, 0)
	charts, err := s.chart_da.GetFavouriteCharts(userId)

	if err != nil {
		return nil, err
	}

	return charts, nil
}

func (s *chartService) GetChartById(chartId string) (models.Chart, error) {
	var chart models.Chart
	chart, err := s.chart_da.GetChartById(chartId)

	return chart, err
}

func (s *chartService) SetFavouriteChart(chartId string) (bool, error) {

	return s.chart_da.SetFavouriteChart(chartId)
}

func (s *chartService) UnsetFavouriteChart(chartId string) (bool, error) {

	return s.chart_da.UnsetFavouriteChart(chartId)
}

func (s *chartService) DeleteChart(chartId string) (bool, error) {

	return s.chart_da.DeleteChart(chartId)
}