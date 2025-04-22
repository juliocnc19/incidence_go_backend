package repository

import (
	"incidence_grade/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Create(role *models.Role) (*models.Role, error) {
	err := r.db.Create(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) FindById(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) Update(role *models.Role) (*models.Role, error) {
	err := r.db.Save(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) Delete(id uint) (map[string]interface{}, error) {
	err := r.db.Delete(&models.Role{}, id).Error
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"ok": true,
		"id": id,
	}
	return result, nil
} 