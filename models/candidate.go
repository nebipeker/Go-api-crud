package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name        string
	Surname     string
	Email       string       `gorm:"unique;not null;type:varchar(100);default:null"`
	JobListings []JobListing `gorm:"many2many"`
}
