package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type fundHandler struct {
	fundRepository repositories.FundInterface
}

func NewFundHandler(fundRepository repositories.FundInterface) *fundHandler {
	return &fundHandler{
		fundRepository,
	}
}

func (fh *fundHandler) CreateFundRequest(c echo.Context) error {
	err := fh.fundRepository.CreateFundRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "fund request created successfully", Status: "success"})
}

func (fh *fundHandler) AcceptFundRequest(c echo.Context) error {
	err := fh.fundRepository.AcceptFundRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "fund request accepted successfully", Status: "success"})
}

func (fh *fundHandler) RejectFundRequest(c echo.Context) error {
	err := fh.fundRepository.RejectFundRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "fund request rejected successfully", Status: "success"})
}
