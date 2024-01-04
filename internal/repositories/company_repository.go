package repositories

import (
	"notification-system/internal/entitys"

	"gorm.io/gorm"
)

type ICompanyRepository interface {
	Create(company *entitys.Companys) (*entitys.Companys, error)
	Get(id int64) (*entitys.Companys, error)
	Update(company *entitys.Companys) (*entitys.Companys, error)
	All() (*entitys.Companys, error)
}

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db}
}

func (r *CompanyRepository) Create(company *entitys.Companys) (*entitys.Companys, error) {
	err := r.db.Create(company).Error
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (r *CompanyRepository) Get(id int64) (*entitys.Companys, error) {
	var company entitys.Companys
	err := r.db.Where("id =?", id).First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Update(company *entitys.Companys) (*entitys.Companys, error) {
	err := r.db.Save(company).Error
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (r *CompanyRepository) All() (*entitys.Companys, error) {
	var company entitys.Companys
	err := r.db.Find(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}
