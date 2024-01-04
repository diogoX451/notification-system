package services

import (
	dtos "notification-system/internal/dtos/company"
	"notification-system/internal/entitys"
	"notification-system/internal/repositories"
)

type ICompanyService interface {
	Create(company dtos.CreateCompanyDTO) (*dtos.CreateMessageDTO, error)
}

type CompanyService struct {
	repository repositories.ICompanyRepository
}

func NewCompanyService(repo repositories.ICompanyRepository) *CompanyService {
	return &CompanyService{repository: repo}
}

func (s *CompanyService) Create(company dtos.CreateCompanyDTO) (*dtos.CreateMessageDTO, error) {
	companyEntity := entitys.Companys{
		Name: company.Name,
	}
	_, err := s.repository.Create(&companyEntity)

	if err != nil {
		return nil, err
	}
	return &dtos.CreateMessageDTO{Message: "Company created"}, nil
}
