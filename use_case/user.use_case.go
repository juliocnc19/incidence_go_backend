package use_case

import (
	"incidence_grade/dto/users"
	"incidence_grade/models"
	"incidence_grade/repository"
)

type User struct {
	repo *repository.UserRespository
}

func NewUser(repo *repository.UserRespository) *User {
	return &User{repo: repo}
}

func (s *User) Create(input dto.CreateUserDto) (*models.User, error) {
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

func (s *User) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *User) GetById(id uint) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *User) Update(id uint, input dto.UpdateUserDto) (*models.User, error) {
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

func (s *User) Delete(id int) (map[string]interface{}, error){
  return s.repo.Detele(id)
}
