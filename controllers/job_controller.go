package controllers

import (
	"net/http"
	"simple-res-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) GetJobListings(ctx echo.Context) error {
	var jobListings []models.JobListing
	result := c.DB.Find(&jobListings)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Joblisting not found")
	}
	return ctx.JSON(http.StatusOK, jobListings)
}
func (c *Controller) GetJobListingById(ctx echo.Context) error {
	var listing models.JobListing
	joblisting_id := ctx.Param("job_id")
	result := c.DB.First(&listing, joblisting_id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Joblisting not found")
	}
	return ctx.JSON(http.StatusOK, listing)
}
func (c *Controller) CreateJobListing(ctx echo.Context) error {
	var joblisting models.JobListing
	if err := ctx.Bind(&joblisting); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}
	c.DB.Create(&joblisting)
	return ctx.JSON(http.StatusOK, joblisting)
}
func (c *Controller) UpdateJobListing(ctx echo.Context) error {
	var joblisting models.JobListing
	joblisting_id := ctx.Param("job_id")
	result := c.DB.First(&joblisting, joblisting_id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Joblisting not found")
	}
	updateData := make(map[string]interface{})
	if err := ctx.Bind(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if title, ok := updateData["Title"]; ok {
		joblisting.Title = title.(string)
	}
	if description, ok := updateData["Description"]; ok {
		joblisting.Description = description.(string)
	}
	if err := c.DB.Save(&joblisting).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update joblisting")
	}
	return ctx.JSON(http.StatusOK, "Company updated succesfully")
}
func (c *Controller) DeleteJobListing(ctx echo.Context) error {
	var joblisting models.JobListing
	joblisting_id := ctx.Param("job_id")
	if err := c.DB.Delete(&joblisting, joblisting_id).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete joblisting")
	}
	return ctx.JSON(http.StatusOK, "Joblisting deleted succesfully")
}
