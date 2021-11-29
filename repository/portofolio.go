package repository

import (
	"web-portfolio-backend/schema"

	"gorm.io/gorm"
)

type PortofolioRepository interface {
	FindByID(ID int) (schema.Portofolio, error)
	FindAll() ([]schema.Portofolio, error)
	Save(about schema.Portofolio) (schema.Portofolio, error)
	Update(portofolio schema.Portofolio) (schema.Portofolio, error)
	DeleteByID(ID int) (bool, error)
}

type portofolioRepository struct {
	db *gorm.DB
}

func NewPortofolioRepository(db *gorm.DB) *portofolioRepository {
	return &portofolioRepository{db}
}

func (r *portofolioRepository) FindByID(ID int) (schema.Portofolio, error) {
	var portfolio schema.Portofolio
	err := r.db.Where("id = ?", ID).Find(&portfolio).Error
	if err != nil {
		return portfolio, err
	}
	return portfolio, nil
}

func (r *portofolioRepository) FindAll() ([]schema.Portofolio, error) {
	var portofolios []schema.Portofolio
	err := r.db.Order("id ASC").Find(&portofolios).Error
	if err != nil {
		return portofolios, err
	}
	return portofolios, nil
}

func (r *portofolioRepository) Save(portofolio schema.Portofolio) (schema.Portofolio, error) {
	err := r.db.Create(&portofolio).Error
	if err != nil {
		return portofolio, err
	}
	return portofolio, nil
}

func (r *portofolioRepository) Update(portofolio schema.Portofolio) (schema.Portofolio, error) {
	err := r.db.Save(&portofolio).Error
	if err != nil {
		return portofolio, err
	}
	return portofolio, nil
}

func (r *portofolioRepository) DeleteByID(ID int) (bool, error) {
	var portofolio schema.About
	err := r.db.Where("id = ?", ID).Delete(&portofolio).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
