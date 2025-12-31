package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type usersHandler struct {
	usersRepository repositories.UsersInterface
}

func NewUsersHandler(usersRepository repositories.UsersInterface) *usersHandler {
	return &usersHandler{
		usersRepository,
	}
}

func (uh *usersHandler) GetAllAdminsRequest(c echo.Context) error {
	res, err := uh.usersRepository.GetAllAdmins(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "admins fetched successfully", Status: "success", Data: map[string]any{"admins": res}})
}

func (uh *usersHandler) GetAllMasterDistributorsRequest(c echo.Context) error {
	res, err := uh.usersRepository.GetAllMasterDistributors(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "master distributors fetched successfully", Status: "success", Data: map[string]any{"mds": res}})
}

func (uh *usersHandler) GetAllDistributorsRequest(c echo.Context) error {
	res, err := uh.usersRepository.GetAllDistributors(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "distributors fetched successfully", Status: "success", Data: map[string]any{"distributors": res}})
}

func (uh *usersHandler) GetAllRetailersRequest(c echo.Context) error {
	res, err := uh.usersRepository.GetAllRetailers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "retailers fetched successfully", Status: "success", Data: map[string]any{"retailers": res}})
}
