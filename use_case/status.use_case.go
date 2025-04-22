package use_case

import (
	dto "incidence_grade/dto/statuses"
	"incidence_grade/models"
	"incidence_grade/repository"
)

type Status struct {
	repo *repository.StatusRepository
}

func NewStatus(repo *repository.StatusRepository) *Status {
	return &Status{repo: repo}
}

func (s *Status) Create(input dto.CreateStatusDto) (*models.Status, error) {
	status := &models.Status{
		Name: input.Name,
	}
	return s.repo.Create(status)
}

func (s *Status) GetAll() ([]models.Status, error) {
	return s.repo.FindAll()
}

func (s *Status) GetById(id uint) (*models.Status, error) {
	return s.repo.FindById(id)
}

func (s *Status) Update(id uint, input dto.UpdateStatusDto) (*models.Status, error) {
	status := &models.Status{
		ID:   id,
		Name: input.Name,
	}
	return s.repo.Update(status)
}

func (s *Status) Delete(id uint) (map[string]interface{}, error) {
	return s.repo.Delete(id)
} 