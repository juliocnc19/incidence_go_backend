package repository

import (
	"incidence_grade/models"
	"gorm.io/gorm"
)

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepository {
	return &StatusRepository{db: db}
}

func (r *StatusRepository) Create(status *models.Status) (*models.Status, error) {
	err := r.db.Create(status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (r *StatusRepository) FindById(id uint) (*models.Status, error) {
	var status models.Status
	err := r.db.First(&status, id).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusRepository) FindAll() ([]models.Status, error) {
	var statuses []models.Status
	err := r.db.Find(&statuses).Error
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

func (r *StatusRepository) Update(status *models.Status) (*models.Status, error) {
	err := r.db.Save(status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (r *StatusRepository) Delete(id uint) (map[string]interface{}, error) {
	err := r.db.Delete(&models.Status{}, id).Error
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"ok": true,
		"id": id,
	}
	return result, nil
} 