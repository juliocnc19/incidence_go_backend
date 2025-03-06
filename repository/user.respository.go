package repository

import (
	"incidence_grade/models"

	"gorm.io/gorm"
)

type UserRespository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRespository {
	return &UserRespository{db: db}
}

func (r *UserRespository) Create(user *models.User) (*models.User, error) {
  error := r.db.Create(user).Error
  if error != nil{
    return nil,error
  }
  return user,nil
}

func (r *UserRespository) FindById(id uint) (*models.User, error) {
	var user models.User
	error := r.db.Preload("Role").First(&user, id).Error
	if error != nil {
		return nil, error
	}
	return &user, nil
}

func (r *UserRespository) FindAll() ([]models.User, error) {
	var users []models.User
	error := r.db.Preload("Role").Find(&users).Error
	return users, error
}

func (r *UserRespository) Update(user *models.User) (*models.User, error) {
  error := r.db.Save(user).Error
  if error != nil{
    return nil,error
  }
  return user,nil
}

func (r *UserRespository) Detele(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
