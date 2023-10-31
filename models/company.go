package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name        string
	Description string
	JobListings []JobListing // One-to-Many relationship might be forced later.
}
