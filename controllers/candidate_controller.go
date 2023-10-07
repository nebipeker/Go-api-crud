package controllers

import (
	"net/http"
	"simple-res-api/models"

	"github.com/labstack/echo/v4"
)

func (c *Controller) GetCandidates(ctx echo.Context) error {
	var candidates []models.Candidate
	c.DB.Find(&candidates)
	return ctx.JSON(http.StatusOK, candidates)
}

func (c *Controller) GetCandidateById(ctx echo.Context) error {
	var candidate models.Candidate
	candidate_id := ctx.Param("candidate_id")
	result := c.DB.First(&candidate, candidate_id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Candidate not found")
	}
	return ctx.JSON(http.StatusOK, candidate)
}

func (c *Controller) CreateCandidate(ctx echo.Context) error {
	var candidate models.Candidate
	if err := ctx.Bind(&candidate); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}
	c.DB.Create(&candidate)
	return ctx.JSON(http.StatusCreated, candidate)
}

func (c *Controller) UpdateCandidate(ctx echo.Context) error {
	candidate_id := ctx.Param("candidate_id")
	var candidate models.Candidate
	result := c.DB.First(&candidate, candidate_id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Candidate not found")
	}

	// Parse the request body into a map
	updateData := make(map[string]interface{})
	if err := ctx.Bind(&updateData); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Check if the "Name" field is present in the request body
	if name, ok := updateData["Name"]; ok {
		candidate.Name = name.(string)
	}

	// Check if the "Surname" field is present in the request body
	if surname, ok := updateData["Surname"]; ok {
		candidate.Surname = surname.(string)
	}

	// Check if the "Surname" field is present in the request body
	if email, ok := updateData["Email"]; ok {
		candidate.Email = email.(string)
	}

	// Update the candidate in the database
	if err := c.DB.Save(&candidate).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update candidate")
	}

	return ctx.JSON(http.StatusOK, "Candidate updated successfully")
}

func (c *Controller) DeleteCandidate(ctx echo.Context) error {
	candidate_id := ctx.Param("candidate_id")
	var candidate models.Candidate
	if err := c.DB.Delete(&candidate, candidate_id).Error; err != nil {
		// Handle the error
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete candidate")
	}

	// If the Delete operation succeeds, it means the candidate was deleted
	return ctx.JSON(http.StatusOK, "Candidate deleted successfully")
}
