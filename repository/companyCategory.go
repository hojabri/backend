package repository

import (
	"fmt"
	"github.com/hojabri/backend/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CompanyCategoryRepository interface {
	Save(companyCategory models.CompanyCategory) (uint, error)
	Update(models.CompanyCategory) error
	Delete(models.CompanyCategory) error
	FindAll() []*models.CompanyCategory
	FindByID(companyCategoryID uint) (*models.CompanyCategory, error)
	DeleteByID(companyCategoryID uint) error
	FindByName(name string) (*models.CompanyCategory, error)
	FindByField(fieldName, fieldValue string) (*models.CompanyCategory, error)
	UpdateSingleField(companyCategory models.CompanyCategory, fieldName, fieldValue string) error
}
type companyCategoryDatabase struct {
	connection *gorm.DB
}

func NewCompanyCategoryRepository() CompanyCategoryRepository {
	if DB == nil {
		_, err = Connect()
		if err != nil {
			log.Error(err)
		}
	}
	return &companyCategoryDatabase{
		connection: DB,
	}
}

func (db companyCategoryDatabase) DeleteByID(companyCategoryID uint) error {
	companyCategory := models.CompanyCategory{}
	companyCategory.ID = companyCategoryID
	result := db.connection.Delete(&companyCategory)
	return result.Error
}

func (db companyCategoryDatabase) Save(companyCategory models.CompanyCategory) (uint, error) {
	result := db.connection.Create(&companyCategory)
	if result.Error != nil {
		return 0, result.Error
	}
	return companyCategory.ID, nil
}

func (db companyCategoryDatabase) Update(companyCategory models.CompanyCategory) error {
	result := db.connection.Save(&companyCategory)
	return result.Error
}

func (db companyCategoryDatabase) Delete(companyCategory models.CompanyCategory) error {
	result := db.connection.Delete(&companyCategory)
	return result.Error
}

func (db companyCategoryDatabase) FindAll() []*models.CompanyCategory {
	var companyCategories []*models.CompanyCategory
	db.connection.Preload(clause.Associations).Find(&companyCategories)
	return companyCategories
}

func (db companyCategoryDatabase) FindByID(companyCategoryID uint) (*models.CompanyCategory, error) {
	var companyCategory models.CompanyCategory
	result := db.connection.Preload(clause.Associations).Find(&companyCategory, "id = ?", companyCategoryID)
	
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &companyCategory, nil
	}
	return nil, nil
}

func (db companyCategoryDatabase) FindByName(name string) (*models.CompanyCategory, error) {
	var companyCategory models.CompanyCategory
	result := db.connection.Preload(clause.Associations).Find(&companyCategory, "name = ?", name)
	
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &companyCategory, nil
	}
	return nil, nil
}

func (db companyCategoryDatabase) FindByField(fieldName, fieldValue string) (*models.CompanyCategory, error) {
	var companyCategory models.CompanyCategory
	result := db.connection.Preload(clause.Associations).Find(&companyCategory, fmt.Sprintf("%s = ?", fieldName), fieldValue)
	
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &companyCategory, nil
	}
	return nil, nil
}

func (db companyCategoryDatabase) UpdateSingleField(companyCategory models.CompanyCategory, fieldName, fieldValue string) error {
	result := db.connection.Model(&companyCategory).Update(fieldName, fieldValue)
	return result.Error
}
