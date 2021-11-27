package service

import (
	"web-portfolio-backend/input"
	"web-portfolio-backend/repository"
	"web-portfolio-backend/schema"
)

type PortfolioService interface {
	PortofolioServiceCreate(input input.InputPortfolio) (schema.Portofolio, error)
	PortofolioServiceGetAll() ([]schema.Portofolio, error)
	PortofolioServiceGetByID(inputID input.InputPortfolioID) (schema.Portofolio, error)
	PortofolioServiceDelete(inputID input.InputPortfolioID) (bool, error)
	PortofolioServiceUpdate(inputID input.InputPortfolioID, inputData input.InputPortfolio, fileLocation string) (schema.Portofolio, error)
}

type portofolioService struct {
	repository repository.PortofolioRepository
}

func NewPortfolioService(repository repository.PortofolioRepository) *portofolioService {
	return &portofolioService{repository}
}

func (s *portofolioService) PortofolioServiceCreate(input input.InputPortfolio) (schema.Portofolio, error) {
	portofolio := schema.Portofolio{}
	portofolio.ID = input.ID
	portofolio.Title = input.Title
	portofolio.ImageCover = input.ImageCover
	portofolio.Phil = input.Phil
	portofolio.Description = input.Description

	newPortfolio, err := s.repository.Save(portofolio)
	if err != nil {
		return newPortfolio, err
	}
	return newPortfolio, nil
}

func (s *portofolioService) PortofolioServiceGetAll() ([]schema.Portofolio, error) {
	portofolios, err := s.repository.FindAll()
	if err != nil {
		return portofolios, err
	}
	return portofolios, nil
}

func (s *portofolioService) PortofolioServiceGetByID(inputID input.InputPortfolioID) (schema.Portofolio, error) {
	portofolio, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return portofolio, err
	}
	return portofolio, nil
}

func (s *portofolioService) PortofolioServiceDelete(inputID input.InputPortfolioID) (bool, error) {
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

func (s *portofolioService) PortofolioServiceUpdate(inputID input.InputPortfolioID, inputData input.InputPortfolio, fileLocation string) (schema.Portofolio, error) {
	portofolio, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return portofolio, err
	}

	portofolio.Title = inputData.Title
	cover := ""
	if fileLocation == "" {
		cover = inputData.ImageCover
	} else {
		cover = fileLocation
	}
	portofolio.ImageCover = cover
	portofolio.Phil = inputData.Phil
	portofolio.Description = inputData.Description
	updatedPorfolio, err := s.repository.Update(portofolio)
	if err != nil {
		return updatedPorfolio, err
	}
	return updatedPorfolio, nil
}
