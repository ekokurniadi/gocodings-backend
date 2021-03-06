package service

import (
	"errors"
	"web-portfolio-backend/input"
	"web-portfolio-backend/repository"
	"web-portfolio-backend/schema"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(input input.LoginInput) (schema.User, error)
	UserServiceGetAll() ([]schema.User, error)
	UserServiceGetByID(ID int) (schema.User, error)
	UserServiceCreate(input input.InputUser) (schema.User, error)
	UserServiceUpdate(inputID input.InputIDUser, inputData input.InputUser, fileLocation string) (schema.User, error)
	UserServiceDelete(inputID input.InputIDUser) (bool, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Login(input input.LoginInput) (schema.User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) UserServiceCreate(input input.InputUser) (schema.User, error) {
	user := schema.User{}
	user.Name = input.Name
	user.Username = input.Username
	// user.Avatar = input.Avatar
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = input.Role

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) UserServiceGetByID(ID int) (schema.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with this ID")
	}

	return user, nil
}

func (s *userService) UserServiceGetAll() ([]schema.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) UserServiceUpdate(inputID input.InputIDUser, inputData input.InputUser, fileLocation string) (schema.User, error) {
	user, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return user, err
	}

	user.Name = inputData.Name
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputData.Password))
	password := ""
	if err != nil {
		password = string(passwordHash)
	} else {
		password = user.Password
	}

	user.Password = password

	userAvatar := ""
	if fileLocation == "" {
		userAvatar = user.Avatar
	} else {
		userAvatar = fileLocation
	}
	user.Avatar = userAvatar
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *userService) UserServiceDelete(inputID input.InputIDUser) (bool, error) {
	_, err := s.repository.DeleteUserByID(inputID.ID)
	if err != nil {
		return false, nil
	}
	return true, nil
}
