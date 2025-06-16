package use_case

import (
	authDto "incidence_grade/dto/auth"
	"incidence_grade/models"
	"incidence_grade/repository"
)

type UserToken struct {
	repo *repository.UserTokenRepository
}

func NewUserToken(repo *repository.UserTokenRepository) *UserToken {
	return &UserToken{repo: repo}
}

func (s *UserToken) SaveDeviceToken(input authDto.DeviceTokenDto) (*models.UserToken, error) {
	existingToken, err := s.repo.FindByToken(input.DeviceToken)
	if err != nil {
		return nil, err
	}

	if existingToken != nil {
		existingToken.UserID = input.UserID
		return s.repo.Create(existingToken)
	}
	userToken := &models.UserToken{
		UserID:      input.UserID,
		DeviceToken: input.DeviceToken,
	}

	return s.repo.Create(userToken)
}

func (s *UserToken) DeleteDeviceToken(id uint) error {
	return s.repo.Delete(id)
}
func (s *UserToken) GetUserDeviceTokens(userID uint) ([]models.UserToken, error) {
	return s.repo.FindByUserID(userID)
}
