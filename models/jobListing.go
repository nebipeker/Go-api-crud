package models

import "gorm.io/gorm"

type JobListing struct {
	gorm.Model
	Title       string
	Description string
	CompanyID   uint
	Candidates  []Candidate `gorm:"many2many:applied_jobs;"`
}
