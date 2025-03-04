package handlers

import (
	"incidence_grade/dto"
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
		Email:     input.Email,
		Username:  input.Username,
		AvatarURL: input.AvatarURL,
		RoleID:    input.RoleID,
	}
	error := s.repo.Create(user)
	if error != nil {
		return nil, error
	}
	return user, nil
}

func (s *UserHandler) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}
