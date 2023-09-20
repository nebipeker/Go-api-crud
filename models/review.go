package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	CandidateId  uint
	JobListingId uint
	CompanyId    uint
	Message      string
}