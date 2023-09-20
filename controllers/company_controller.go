package controllers

import (
	"net/http"
	"simple-res-api/models"

	"github.com/labstack/echo/v4"
)

func (c *Controller) GetCompanies(ctx echo.Context) error {
	var companies []models.Company
	result := c.DB.Find(&companies)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Companies not found")
	}
	return ctx.JSON(http.StatusOK, companies)

}

func (c *Controller) GetCompanyById(ctx echo.Context) error {
	var company models.Company
	company_id := ctx.Param("company_id")
	result := c.DB.First(&company, company_id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Company not found")
	}
	return ctx.JSON(http.StatusOK, company)
}

func (c *Controller) CreateCompany(ctx echo.Context) error {
	var company models.Company
	if err := ctx.Bind(&company); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}
	c.DB.Create(&company)
	return ctx.JSON(http.StatusCreated, company)
}

func (c *Controller) UpdateCompany(ctx echo.Context) error {
	var company models.Company
	company_id := ctx.Param("company_id")
	result := c.DB.First(&company, company_id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Company not found")
	}

	// Parse the request body into a map
	updateData := make(map[string]interface{})
	if err := ctx.Bind(&updateData); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Check if the "Name" field is present in the request body
	if name, ok := updateData["Name"]; ok {
		company.Name = name.(string)
	}

	// Check if the "Description" field is present in the request body
	if description, ok := updateData["Description"]; ok {
		company.Description = description.(string)
	}

	if err := c.DB.Save(&company).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update company")
	}

	return ctx.JSON(http.StatusOK, "Company updated successfully")
}

func (c *Controller) DeleteCompany(ctx echo.Context) error {
	var company models.Company
	company_id := ctx.Param("company_id")
	if err := c.DB.Delete(&company, company_id).Error; err != nil {
		// Handle the error
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete company")
	}
	return ctx.JSON(http.StatusOK, "Company deleted successfully")
}
