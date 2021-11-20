package service

import (
	"web-portfolio-backend/input"
	"web-portfolio-backend/repository"
	"web-portfolio-backend/schema"
)

type AboutService interface {
	AboutServiceGetAll() ([]schema.About, error)
	AboutServiceGetByID(inputID input.InputID) (schema.About, error)
	AboutServiceCreate(input input.InputAbout) (schema.About, error)
	AboutServiceUpdate(inputID input.InputID, inputData input.InputAbout) (schema.About, error)
	AboutServiceDelete(inputID input.InputID) (bool, error)
}

type aboutService struct {
	repository repository.AboutRepository
}

func NewAboutService(repository repository.AboutRepository) *aboutService {
	return &aboutService{repository}
}

func (s *aboutService) AboutServiceCreate(input input.InputAbout) (schema.About, error) {
	about := schema.About{}
	about.TentangSaya = input.TentangSaya
	about.Alamat = input.Alamat
	about.Telp = input.Telp
	about.Whatsapp = input.Whatsapp
	newAbout, err := s.repository.Save(about)
	if err != nil {
		return newAbout, err
	}
	return newAbout, nil
}
