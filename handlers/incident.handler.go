package handlers

import (
	dto "incidence_grade/dto/incidents"
	"incidence_grade/models"
	"incidence_grade/repository"
)

type IncidentHandler struct{
  repo *repository.IncidentRepository
}

func NewIncidentHandler(repo *repository.IncidentRepository) *IncidentHandler{
  return &IncidentHandler{repo:repo}
}

func (s *IncidentHandler) CreateIncident(input dto.CreateIncidentDto) (*models.Incident,error){
  incident := &models.Incident{
    Title: input.Title,
    Description: input.Description,
    StatusID: input.StatusID,
    Response: input.Response,
    UserID: input.UserID,
    AttachmentPath: input.AttachmentPath,
  }
  return s.repo.Create(incident)
}

func (s *IncidentHandler) GetAllIncidents() ([]models.Incident,error){
  return s.repo.FindAll()
}

func (s *IncidentHandler) GetIncidentById(id uint) (*models.Incident, error){
  return s.repo.FindById(id)
} 

func (s *IncidentHandler) UpdateIncident(id uint, input dto.UpdateIncidentDto) (*models.Incident,error){
  incidentUpdate := &models.Incident{
    ID: id,
    Title: input.Title,
    Description: input.Description,
    StatusID: input.StatusID,
    Response: input.Response,
    UserID: input.UserID,
    AttachmentPath: input.AttachmentPath,
  }
  return s.repo.Update(incidentUpdate)
}

func (s *IncidentHandler) DeleteIncident(id uint) (map[string]interface{},error){
  return s.repo.Delete(id)
}
