package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name         string          `json:"name" gorm:"uniqueIndex"`
	Description  string          `json:"description,omitempty"`
	CategoryID   *uint            `json:"category_id,omitempty"`
	Category     *CompanyCategory `json:"category"`
	Website      string          `json:"website,omitempty"`
}

type CompanyCategory struct {
	gorm.Model
	Name string `json:"name" gorm:"uniqueIndex"`
}
