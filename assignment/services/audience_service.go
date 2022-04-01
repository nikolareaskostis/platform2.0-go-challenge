package services

import (
	"GoRepo/source/assignment/dataAccess"
	"GoRepo/source/assignment/models"
)

type AudienceService interface {
	GetAllAudiences(userId string) ([]models.Audience, error)
	GetFavouriteAudiences(userId string) ([]models.Audience, error)
	GetAudienceById(userId string) (models.Audience, error)
	SetFavouriteAudience(audienceId string) (bool, error)
	UnsetFavouriteAudience(audienceId string) (bool, error)
	DeleteAudience(audienceId string) (bool, error)
}

type audienceService struct {
	audience_da dataAccess.Audience_da
}

func NewAudienceService(audience_da dataAccess.Audience_da) AudienceService {
	return &audienceService{audience_da: audience_da}
}

func (s *audienceService) GetAllAudiences(userId string) ([]models.Audience, error) {
	audiences := make([]models.Audience, 0)
	audiences, err := s.audience_da.GetAllAudiences(userId)

	if err != nil {
		return nil, err
	}

	return audiences, nil
}

func (s *audienceService) GetFavouriteAudiences(userId string) ([]models.Audience, error) {
	audiences := make([]models.Audience, 0)
	audiences, err := s.audience_da.GetFavouriteAudiences(userId)

	if err != nil {
		return nil, err
	}

	return audiences, nil
}

func (s *audienceService) GetAudienceById(audienceId string) (models.Audience, error) {
	var audience models.Audience
	audience, err := s.audience_da.GetAudienceById(audienceId)

	return audience, err
}

func (s *audienceService) SetFavouriteAudience(audienceId string) (bool, error) {

	return s.audience_da.SetFavouriteAudience(audienceId)
}

func (s *audienceService) UnsetFavouriteAudience(audienceId string) (bool, error) {

	return s.audience_da.UnsetFavouriteAudience(audienceId)
}

func (s *audienceService) DeleteAudience(audienceId string) (bool, error) {

	return s.audience_da.DeleteAudience(audienceId)
}
