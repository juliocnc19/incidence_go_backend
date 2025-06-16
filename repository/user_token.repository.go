package repository

import (
	"errors"
	"incidence_grade/models"
	"gorm.io/gorm"
)

type UserTokenRepository struct {
	db *gorm.DB
}

func NewUserTokenRepository(db *gorm.DB) *UserTokenRepository {
	return &UserTokenRepository{db: db}
}

func (r *UserTokenRepository) Create(userToken *models.UserToken) (*models.UserToken, error) {
	err := r.db.Create(userToken).Error
	if err != nil {
		return nil, err
	}
	return userToken, nil
}

func (r *UserTokenRepository) Delete(id uint) error {
	result := r.db.Delete(&models.UserToken{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("token no encontrado")
	}
	return nil
}

func (r *UserTokenRepository) FindByUserID(userID uint) ([]models.UserToken, error) {
	var tokens []models.UserToken
	err := r.db.Where("user_id = ?", userID).Find(&tokens).Error
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *UserTokenRepository) FindByToken(token string) (*models.UserToken, error) {
	var userToken models.UserToken
	err := r.db.Where("device_token = ?", token).First(&userToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userToken, nil
}
