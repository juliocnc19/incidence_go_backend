package repository

import (
	"errors"
	"incidence_grade/models"
	"incidence_grade/utils"

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
	if error != nil {
		return nil, error
	}
	return user, nil
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
	if error != nil {
		return nil, error
	}
	return users, nil
}

func (r *UserRespository) Update(user *models.User) (*models.User, error) {
	if err := r.db.
		Model(&models.User{ID: user.ID}).
		Updates(user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRespository) Detele(id int) (map[string]interface{}, error) {
	error := r.db.Delete(&models.User{}, id).Error
	if error != nil {
		return nil, error
	}
	resutl := map[string]interface{}{
		"ok": true,
		"id": id,
	}
	return resutl, nil
}

func (r *UserRespository) Login(email string, password string) (*models.User, error) {
	var user models.User
	error := r.db.Where("email = ?", email).Preload("Role").First(&user).Error
	if error != nil {
		return nil, errors.New("usuario no encontrado")
	}

	result := utils.CheckPasswordHash(password, user.Password)
	if !result {
		return nil, errors.New("contraseña invalida")
	}

	return &user, nil
}

// UpdatePassword actualiza la contraseña de un usuario
func (r *UserRespository) UpdatePassword(userID uint, hashedPassword string) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("password", hashedPassword).Error
}
