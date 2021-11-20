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
	about.Email = input.Email
	newAbout, err := s.repository.Save(about)
	if err != nil {
		return newAbout, err
	}
	return newAbout, nil
}

func (s *aboutService) AboutServiceUpdate(inputID input.InputID, inputData input.InputAbout) (schema.About, error) {
	about, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return about, err
	}
	about.TentangSaya = inputData.TentangSaya
	about.Alamat = inputData.Alamat
	about.Telp = inputData.Telp
	about.Whatsapp = inputData.Whatsapp

	updatedAbout, err := s.repository.Update(about)
	if err != nil {
		return updatedAbout, err
	}

	return updatedAbout, nil
}

func (s *aboutService) AboutServiceGetAll() ([]schema.About, error) {
	abouts, err := s.repository.FindAll()
	if err != nil {
		return abouts, err
	}
	return abouts, nil
}

func (s *aboutService) AboutServiceGetByID(inputID input.InputID) (schema.About, error) {
	about, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return about, err
	}
	return about, nil
}

func (s *aboutService) AboutServiceDelete(inputID input.InputID) (bool, error) {
	_, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return false, err
	}
	_, err = s.repository.DeleteByID(inputID.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
