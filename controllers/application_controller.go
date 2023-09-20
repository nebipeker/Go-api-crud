package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func (c *Controller) ApplyForJob(ctx echo.Context) error {
	var requestData map[string]interface{}
	//job_id := ctx.Param("job_id")

	if err := ctx.Bind(&requestData); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	return ctx.JSON(http.StatusOK, "")
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
