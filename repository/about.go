package repository

import (
	"web-portfolio-backend/schema"

	"gorm.io/gorm"
)

type AboutRepository interface {
	Save(about schema.About) (schema.About, error)
	FindByID(ID int) (schema.About, error)
	FindAll() ([]schema.About, error)
	Update(about schema.About) (schema.About, error)
	DeleteByID(ID int) (schema.About, error)
}

type aboutRepository struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) *aboutRepository {
	return &aboutRepository{db}
}

func (r *aboutRepository) FindByID(ID int) (schema.About, error) {
	var about schema.About
	err := r.db.Where("id = ?", ID).Find(&about).Error
	if err != nil {
		return about, err
	}
	return about, nil
}

func (r *aboutRepository) FindAll() ([]schema.About, error) {
	var abouts []schema.About
	err := r.db.Find(&abouts).Error
	if err != nil {
		return abouts, err
	}
	return abouts, nil
}

func (r *aboutRepository) Save(about schema.About) (schema.About, error) {
	err := r.db.Create(&about).Error
	if err != nil {
		return about, err
	}
	return about, nil
}

func (r *aboutRepository) Update(about schema.About) (schema.About, error) {
	err := r.db.Save(&about).Error
	if err != nil {
		return about, err
	}
	return about, nil
}

func (r *aboutRepository) DeleteByID(ID int) (schema.About, error) {
	var about schema.About
	err := r.db.Where("id = ?", ID).Delete(&about).Error
	if err != nil {
		return about, err
	}
	return about, nil
}
