package repository

import (
	"gorm.io/gorm"
	"incidence_grade/models"
)

type AttachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) *AttachmentRepository {
	return &AttachmentRepository{db: db}
}

func (r *AttachmentRepository) Create(attachment *models.Attachment) (*models.Attachment, error) {
	err := r.db.Create(attachment).Error
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (r *AttachmentRepository) FindById(id uint) (*models.Attachment, error) {
	var attachment models.Attachment
	err := r.db.Preload("Incident").First(&attachment, id).Error
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *AttachmentRepository) FindAll() ([]models.Attachment, error) {
	var attachments []models.Attachment
	err := r.db.Preload("Incident").Find(&attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (r *AttachmentRepository) Update(attachment *models.Attachment) (*models.Attachment, error) {
	err := r.db.Save(attachment).Error
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (r *AttachmentRepository) Delete(id uint) (map[string]interface{}, error) {
	err := r.db.Delete(&models.Attachment{}, id).Error
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"ok": true,
		"id": id,
	}
	return result, nil
}

func (r *AttachmentRepository) FindByIncidentId(incidentId uint) ([]models.Attachment, error) {
	var attachments []models.Attachment
	err := r.db.Where("incident_id = ?", incidentId).Find(&attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

