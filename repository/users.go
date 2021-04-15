package repository

import (
	"fmt"
	"github.com/hojabri/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

)

type UserRepository interface {
	Save(user models.User) (int, error)
	Update(models.User) error
	Delete(models.User) error
	FindAll() []*models.User
	FindByID(userID int) (*models.User, error)
	DeleteByID(userID int) error
	FindByName(name string) (*models.User, error)
	FindByField(fieldName, fieldValue string) (*models.User, error)
	UpdateSingleField(user models.User, fieldName, fieldValue string) error
}
type userDatabase struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	if DB == nil {
		Connect()
	}
	return &userDatabase{
		connection: DB,
	}
}

func (db userDatabase) DeleteByID(userID int) error {
	user := models.User{}
	user.ID = userID
	result := db.connection.Delete(&user)
	return result.Error
}

func (db userDatabase) Save(user models.User) (int, error) {
	result := db.connection.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (db userDatabase) Update(user models.User) error {
	result := db.connection.Save(&user)
	return result.Error
}

func (db userDatabase) Delete(user models.User) error {
	result := db.connection.Delete(&user)
	return result.Error
}

func (db userDatabase) FindAll() []*models.User {
	var users []*models.User
	db.connection.Preload(clause.Associations).Find(&users)
	return users
}

func (db userDatabase) FindByID(userID int) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, "id = ?", userID)
	
	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}


func (db userDatabase) FindByName(name string) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, "name = ?", name)
	
	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) FindByField(fieldName, fieldValue string) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, fmt.Sprintf("%s = ?", fieldName) ,fieldValue)
	
	if result.Error!=nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) UpdateSingleField(user models.User, fieldName, fieldValue string) error {
	result := db.connection.Model(&user).Update(fieldName, fieldValue)
	return result.Error
}