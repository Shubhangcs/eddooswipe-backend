package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type commisionHandler struct {
	commisionRepository repositories.CommisionInterface
}

func NewCommisionHandler(commisionRepository repositories.CommisionInterface) *commisionHandler {
	return &commisionHandler{
		commisionRepository,
	}
}

func (ch *commisionHandler) CreateCommisionRequest(c echo.Context) error {
	err := ch.commisionRepository.CreateCommision(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "commision created successfully", Status: "success"})
}

func (ch *commisionHandler) GetAllCommisionsRequest(c echo.Context) error {
	res, err := ch.commisionRepository.GetAllCommisions(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "commisions fetched successfully", Status: "success", Data: map[string]any{"commisions": res}})
}

func (ch *commisionHandler) UpdateCommisionRequest(c echo.Context) error {
	err := ch.commisionRepository.UpdateCommision(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "commision updated successfully", Status: "success"})
}

func (ch *commisionHandler) DeleteCommisionRequest(c echo.Context) error {
	err := ch.commisionRepository.DeleteCommision(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "commision delete successfully", Status: "success"})
}
