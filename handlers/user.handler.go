package handlers

import (
	"incidence_grade/dto/users"
	"incidence_grade/models"
	"incidence_grade/repository"
)

type UserHandler struct {
	repo *repository.UserRespository
}

func NewUserHandler(repo *repository.UserRespository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (s *UserHandler) CreateUser(input dto.CreateUserDto) (*models.User, error) {
	user := &models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
		Email:     input.Email,
		Username:  input.Username,
		AvatarURL: input.AvatarURL,
		RoleID:    input.RoleID,
	}
	return s.repo.Create(user)
}

func (s *UserHandler) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserHandler) GetUserById(id uint) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *UserHandler) UpdateUser(id uint, input dto.UpdateUserDto) (*models.User, error) {
	user := &models.User{
		ID:        id,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		Username:  input.Username,
		AvatarURL: input.AvatarURL,
		RoleID:    input.RoleID,
	}
	return s.repo.Update(user)
}

func (s *UserHandler) DeleteUser(id int) (map[string]interface{}, error){
  return s.repo.Detele(id)
}
