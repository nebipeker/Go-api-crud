package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Review struct {
	DB *gorm.DB
}

func (c *Controller) ListJobReviews(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
func (c *Controller) CreateJobReview(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
func (c *Controller) ListCandidateReviews(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
func (c *Controller) CreateCandidateReview(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}
