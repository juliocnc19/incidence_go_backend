package use_case

import (
	authDto "incidence_grade/dto/auth"
	dto "incidence_grade/dto/users"
	"incidence_grade/models"
	"incidence_grade/repository"
	"incidence_grade/utils"
)

type User struct {
	repo *repository.UserRespository
}

func NewUser(repo *repository.UserRespository) *User {
	return &User{repo: repo}
}

func (s *User) Create(input dto.CreateUserDto) (*models.User, error) {
	hashedPassword, error := utils.HashPassword(input.Password)
	if error != nil {
		panic("Error al hashear las contraseña")
	}
	user := &models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  hashedPassword,
		Email:     input.Email,
		Username:  input.Username,
		RoleID:    input.RoleID,
		Cedula:    input.Cedula,
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
	hashedPassword, error := utils.HashPassword(input.Password)
	if error != nil {
		panic("Error al hashear las contraseña")
	}

	user := &models.User{
		ID:        id,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
		Username:  input.Username,
		RoleID:    input.RoleID,
	}
	return s.repo.Update(user)
}

func (s *User) Delete(id int) (map[string]interface{}, error) {
	return s.repo.Detele(id)
}

func (s *User) Login(input authDto.LoginUserDto) (*models.User, error) {
	return s.repo.Login(input.Email, input.Password)
}

func (s *User) Register(input authDto.RegisterUserDto) (*models.User, error) {
	hashedPassword, error := utils.HashPassword(input.Password)
	if error != nil {
		panic("Error al hashear las contraseña")
	}
	user := &models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
		Username:  input.Username,
		RoleID:    2,
	}
	return s.repo.Create(user)
}
