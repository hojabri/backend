package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email,omitempty" gorm:"uniqueIndex"`
	Phone     string  `json:"phone,omitempty"`
	CompanyID *uint    `json:"company_id,omitempty"`
	Company   *Company `json:"company,omitempty"`
}
