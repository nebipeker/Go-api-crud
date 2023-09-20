package models

import "gorm.io/gorm"

type JobListing struct {
	gorm.Model
	Title       string
	Description string
	CompanyID   uint        // Foreign key
	Candidates  []Candidate `gorm:"many2many:job_listing_candidates;"`
}
