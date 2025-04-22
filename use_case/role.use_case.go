package use_case

import (
	dto "incidence_grade/dto/roles"
	"incidence_grade/models"
	"incidence_grade/repository"
)

type Role struct {
	repo *repository.RoleRepository
}

func NewRole(repo *repository.RoleRepository) *Role {
	return &Role{repo: repo}
}

func (s *Role) Create(input dto.CreateRoleDto) (*models.Role, error) {
	role := &models.Role{
		Name: input.Name,
	}
	return s.repo.Create(role)
}

func (s *Role) GetAll() ([]models.Role, error) {
	return s.repo.FindAll()
}

func (s *Role) GetById(id uint) (*models.Role, error) {
	return s.repo.FindById(id)
}

func (s *Role) Update(id uint, input dto.UpdateRoleDto) (*models.Role, error) {
	role := &models.Role{
		ID:   id,
		Name: input.Name,
	}
	return s.repo.Update(role)
}

func (s *Role) Delete(id uint) (map[string]interface{}, error) {
	return s.repo.Delete(id)
} 