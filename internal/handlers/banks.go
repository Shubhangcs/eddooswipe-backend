package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type bankHandler struct {
	bankRepository repositories.BankInterface
}

func NewBankHandler(bankRepository repositories.BankInterface) *bankHandler {
	return &bankHandler{
		bankRepository,
	}
}

func (bh *bankHandler) CreateAdminBankRequest(c echo.Context) error {
	err := bh.bankRepository.CreateAdminBank(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "admin bank created successfully", Status: "success"})
}

func (bh *bankHandler) CreateRetailerBankRequest(c echo.Context) error {
	err := bh.bankRepository.CreateRetailerBank(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "retailer bank created successfully", Status: "success"})
}

func (bh *bankHandler) GetAdminBanksByAdminIDRequest(c echo.Context) error {
	res, err := bh.bankRepository.GetAdminBanksByAdminID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "banks fetched successfully", Status: "success", Data: map[string]any{"banks": res}})
}

func (bh *bankHandler) GetRetailerBanksByRetailerIDRequest(c echo.Context) error {
	res, err := bh.bankRepository.GetRetailerBanksByRetailerID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "banks fetched successfully", Status: "success", Data: map[string]any{"banks": res}})
}

func (bh *bankHandler) GetAllAdminBanksRequest(c echo.Context) error {
	res, err := bh.bankRepository.GetAllAdminBanks(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "banks fetched successfully", Status: "success", Data: map[string]any{"banks": res}})
}

func (bh *bankHandler) GetAllRetailerBanksRequest(c echo.Context) error {
	res, err := bh.bankRepository.GetAllRetailerBanks(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "banks fetched successfully", Status: "success", Data: map[string]any{"banks": res}})
}

func (bh *bankHandler) DeleteAdminBankRequest(c echo.Context) error {
	err := bh.bankRepository.DeleteAdminBank(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "admin bank deleted successfully", Status: "success"})
}

func (bh *bankHandler) DeleteRetailerBankRequest(c echo.Context) error {
	err := bh.bankRepository.DeleteRetailerBank(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "retailer bank deleted successfully", Status: "success"})
}