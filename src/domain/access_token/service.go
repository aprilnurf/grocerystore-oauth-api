package access_token

import (
	"github.com/aprilnurf/grocerystore-oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpiredTime(AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpiredTime(AccessToken) *errors.RestError
}

type service struct {
	repository Repository
}

func (s *service) Create(token AccessToken) *errors.RestError {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.Create(token)
}

func (s *service) UpdateExpiredTime(token AccessToken) *errors.RestError {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpiredTime(token)
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(token string) (*AccessToken, *errors.RestError) {
	accessTokenId := strings.TrimSpace(token)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token")
	}
	accessToken, err := s.repository.GetById(token)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
