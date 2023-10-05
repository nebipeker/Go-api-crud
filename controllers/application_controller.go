package controllers

import (
	"net/http"
	"simple-res-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func (c *Controller) ApplyForJob(ctx echo.Context) error {
	candidateID := ctx.Param("candidateID")
	jobListingID := ctx.Param("jobListingID")

	// Check if the candidate exists
	var candidate models.Candidate
	if err := c.DB.First(&candidate, candidateID).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Candidate not found"})
	}

	// Check if the job listing exists
	var jobListing models.JobListing
	if err := c.DB.First(&jobListing, jobListingID).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Job listing not found"})
	}

	// Append the job listing to the candidate's list of applied jobs
	c.DB.Model(&candidate).Association("JobListings").Append(&jobListing)

	return ctx.JSON(http.StatusCreated, map[string]string{"message": "Application submitted"})

}
func (c *Controller) ListApplications(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
func (c *Controller) GetApplication(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
func (c *Controller) UpdateApplication(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
func (c *Controller) WithdrawApplication(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
