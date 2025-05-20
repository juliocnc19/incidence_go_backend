package use_case

import (
	dto "incidence_grade/dto/incidents"
	"incidence_grade/models"
	"incidence_grade/repository"
	"log"
)

const (
	StatusInProgress = 1
	StatusDraft      = 4
)

type Incident struct {
	repo *repository.IncidentRepository
}

func NewIncident(repo *repository.IncidentRepository) *Incident {
	return &Incident{repo: repo}
}

func (s *Incident) Create(input dto.CreateIncidentDto) (*models.Incident, error) {
	incident := &models.Incident{
		Title:       input.Title,
		Description: input.Description,
		StatusID:    input.StatusID,
		Response:    input.Response,
		UserID:      input.UserID,
	}
	return s.repo.Create(incident)
}

func (s *Incident) GetAll() ([]models.Incident, error) {
	return s.repo.FindAll()
}

func (s *Incident) GetById(id uint) (*models.Incident, error) {
	return s.repo.FindById(id)
}

func (s *Incident) Update(id uint, input dto.UpdateIncidentDto) (*models.Incident, error) {
	incidentUpdate := &models.Incident{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		StatusID:    input.StatusID,
		Response:    input.Response,
		UserID:      input.UserID,
	}
	return s.repo.Update(incidentUpdate)
}

func (s *Incident) Delete(id uint) (map[string]interface{}, error) {
	return s.repo.Delete(id)
}

func (s *Incident) FindByIdUser(user_id uint) ([]models.Incident, error) {
	return s.repo.FindByIdUser(user_id)
}

func (s *Incident) SaveFiles(filenames []string, incident_id uint) ([]models.Attachment, error) {
	var uploadedAttachments []models.Attachment
	for file := range filenames {
		attachment := models.Attachment{
			AttachmentPath: filenames[file],
			IncidentID:     incident_id,
		}
		uploadedAttachments = append(uploadedAttachments, attachment)
	}

	// Save the files
	attachments, err := s.repo.SaveFile(uploadedAttachments)
	if err != nil {
		return nil, err
	}

	// Update incident status from draf to in_progress
	incident, err := s.GetById(incident_id)
	if err != nil {
		log.Printf("Error getting incident %d: %v", incident_id, err)
		return attachments, nil // Return attachments even if status update fails
	}

	if incident == nil {
		log.Printf("Incident %d not found", incident_id)
		return attachments, nil
	}

	if incident.StatusID == StatusDraft {
		log.Printf("Updating incident %d status from draft to in_progress", incident_id)
		updateDto := dto.UpdateIncidentDto{
			Title:       incident.Title,
			Description: incident.Description,
			StatusID:    StatusInProgress,
			Response:    incident.Response,
			UserID:      incident.UserID,
		}
		_, err = s.Update(incident_id, updateDto)
		if err != nil {
			log.Printf("Error updating incident %d status: %v", incident_id, err)
			return attachments, nil // Return attachments even if status update fails
		}
		log.Printf("Successfully updated incident %d status to in_progress", incident_id)
	}

	return attachments, nil
}
