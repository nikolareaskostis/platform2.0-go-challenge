package services

import (
	"GoRepo/source/assignment/dataAccess"
	"GoRepo/source/assignment/models"
)

type InsightService interface {
	GetAllInsights(userId string) ([]models.Insight, error)
	GetFavouriteInsights(userId string) ([]models.Insight, error)
	GetInsightById(insightId string) (models.Insight, error)
	SetFavouriteInsight(insightId string) (bool, error)
	UnsetFavouriteInsight(insightId string) (bool, error)
	DeleteInsight(chartId string) (bool, error)
}

type insightService struct {
	insight_da dataAccess.Insight_da
}

func NewInsightService(insight_da dataAccess.Insight_da) InsightService {
	return &insightService{insight_da: insight_da}
}

func (s *insightService) GetAllInsights(userId string) ([]models.Insight, error) {
	insights := make([]models.Insight, 0)
	insights, err := s.insight_da.GetAllInsights(userId)

	if err != nil {
		return nil, err
	}

	return insights, nil
}

func (s *insightService) GetFavouriteInsights(userId string) ([]models.Insight, error) {
	insights := make([]models.Insight, 0)
	insights, err := s.insight_da.GetFavouriteInsights(userId)

	if err != nil {
		return nil, err
	}

	return insights, nil
}

func (s *insightService) GetInsightById(insightId string) (models.Insight, error) {
	var insight models.Insight
	insight, err := s.insight_da.GetInsightById(insightId)

	return insight, err
}

func (s *insightService) SetFavouriteInsight(insightId string) (bool, error) {

	return s.insight_da.SetFavouriteInsight(insightId)
}

func (s *insightService) UnsetFavouriteInsight(insightId string) (bool, error) {

	return s.insight_da.UnsetFavouriteInsight(insightId)
}

func (s *insightService) DeleteInsight(insightId string) (bool, error) {

	return s.insight_da.DeleteInsight(insightId)
}