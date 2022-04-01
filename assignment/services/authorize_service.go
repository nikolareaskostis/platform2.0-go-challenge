package services

import (
	"GoRepo/source/assignment/dataAccess"
	"GoRepo/source/assignment/models"
)

type AuthorizeService interface {
	Login(user models.User) (bool, error)
	Register(user models.User) (bool, error)
	Check(user models.User) (bool, error)
}

type authorizeService struct {
	authorize_da dataAccess.Authorize_da
}

func NewAuthService(authorize_da dataAccess.Authorize_da) AuthorizeService {
	return &authorizeService{authorize_da: authorize_da}
}

func (s *authorizeService) Login(user models.User) (bool, error) {

	return s.authorize_da.Login(user)
}

func (s *authorizeService) Register(user models.User) (bool, error) {

	return s.authorize_da.Register(user)
}

func (s *authorizeService) Check(user models.User) (bool, error) {

	return s.authorize_da.Check(user)
}