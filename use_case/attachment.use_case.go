package use_case

import (
	dto "incidence_grade/dto/attachments"
	"incidence_grade/models"
	"incidence_grade/repository"
)

type Attachment struct {
	repo *repository.AttachmentRepository
}

func NewAttachment(repo *repository.AttachmentRepository) *Attachment {
	return &Attachment{repo: repo}
}

func (s *Attachment) Create(input dto.CreateAttachmentDto) (*models.Attachment, error) {
	attachment := &models.Attachment{
		AttachmentPath: input.AttachmentPath,
		IncidentID:     input.IncidentID,
	}
	return s.repo.Create(attachment)
}

func (s *Attachment) GetAll() ([]models.Attachment, error) {
	return s.repo.FindAll()
}

func (s *Attachment) GetById(id uint) (*models.Attachment, error) {
	return s.repo.FindById(id)
}

func (s *Attachment) Update(id uint, input dto.UpdateAttachmentDto) (*models.Attachment, error) {
	attachment := &models.Attachment{
		ID:             id,
		AttachmentPath: input.AttachmentPath,
		IncidentID:     input.IncidentID,
	}
	return s.repo.Update(attachment)
}

func (s *Attachment) Delete(id uint) (map[string]interface{}, error) {
	return s.repo.Delete(id)
}

func (s *Attachment) GetByIncidentId(incidentId uint) ([]models.Attachment, error) {
	return s.repo.FindByIncidentId(incidentId)
} 