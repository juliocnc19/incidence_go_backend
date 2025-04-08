package repository

import (
	"incidence_grade/models"

	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{db: db}
}

func (r *IncidentRepository) Create(incident *models.Incident) (*models.Incident, error) {
	error := r.db.Create(incident).Error
	if error != nil {
		return nil, error
	}
	return incident, nil
}

func (r *IncidentRepository) FindById(id uint) (*models.Incident, error) {
	var incident models.Incident
	error := r.db.Preload("Status").Preload("User").First(&incident,id).Error
	if error != nil {
		return nil, error
	}
	return &incident, nil
}

func (r *IncidentRepository) FindAll() ([]models.Incident, error) {
	var incidents []models.Incident
	error := r.db.Preload("Status").Preload("User").Find(&incidents).Error
	if error != nil {
		return nil, error
	}
	return incidents, nil
}

func (r *IncidentRepository) Update(incident *models.Incident) (*models.Incident, error) {
	error := r.db.Save(incident).Error
	if error != nil {
		return nil, error
	}
	return incident, nil
}

func (r *IncidentRepository) Delete(id uint) (map[string]interface{}, error) {

	error := r.db.Delete(&models.Incident{}, id).Error
	if error != nil {
		return nil, error
	}
	resutl := map[string]interface{}{
		"ok":      true,
		"id":      id,
	}
	return resutl, nil
}

func (r *IncidentRepository) FindByIdUser(user_id uint) ([]models.Incident, error){
	var incident []models.Incident
	error := r.db.Where("user_id = ?",user_id).Preload("Status").Preload("User").Preload("Attachment").Find(&incident).Error
	if error != nil {
		return nil, error
	}
	return incident, nil
}
