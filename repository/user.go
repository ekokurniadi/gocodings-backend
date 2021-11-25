package repository

import (
	"web-portfolio-backend/schema"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user schema.User) (schema.User, error)
	FindByUsername(username string) (schema.User, error)
	FindByID(ID int) (schema.User, error)
	Update(user schema.User) (schema.User, error)
	FindAll() ([]schema.User, error)
	DeleteUserByID(ID int) (schema.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user schema.User) (schema.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByUsername(username string) (schema.User, error) {
	var user schema.User
	err := r.db.Where("username = ?", username).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID int) (schema.User, error) {
	var user schema.User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *userRepository) Update(user schema.User) (schema.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindAll() ([]schema.User, error) {
	var users []schema.User
	err := r.db.Order("id desc").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) DeleteUserByID(ID int) (schema.User, error) {
	var user schema.User
	err := r.db.Where("id = ?", ID).Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
