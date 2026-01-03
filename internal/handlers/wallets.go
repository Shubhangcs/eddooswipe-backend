package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type walletHandler struct {
	walletRepository repositories.WalletInterface
}

func NewWalletHandler(walletRepository repositories.WalletInterface) *walletHandler {
	return &walletHandler{
		walletRepository,
	}
}

func (wh *walletHandler) GetAdminWalletBalanceRequest(c echo.Context) error {
	res, err := wh.walletRepository.GetAdminWalletBalance(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "admin wallet balance fetched successfully", Status: "success", Data: map[string]any{"balance": res}})
}

func (wh *walletHandler) GetMasterDistributorWalletBalanceRequest(c echo.Context) error {
	res, err := wh.walletRepository.GetMasterDistributorWalletBalance(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "master distributor wallet balance fetched successfully", Status: "success", Data: map[string]any{"balance": res}})
}

func (wh *walletHandler) GetDistributorWalletBalanceRequest(c echo.Context) error {
	res, err := wh.walletRepository.GetDistributorWalletBalance(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "distributor wallet balance fetched successfully", Status: "success", Data: map[string]any{"balance": res}})
}

func (wh *walletHandler) GetRetailerWalletBalanceRequest(c echo.Context) error {
	res, err := wh.walletRepository.GetRetailerWalletBalance(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "retailer wallet balance fetched successfully", Status: "success", Data: map[string]any{"balance": res}})
}

func (wh *walletHandler) AdminWalletTopupRequest(c echo.Context) error {
	err := wh.walletRepository.AdminWalletTopup(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "admin wallet topup success", Status: "success"})
}

func (wh *walletHandler) GetLedgerTransactionsRequest(c echo.Context) error {
	res, err := wh.walletRepository.GetLedgerTransactions(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "ledger entries fetched successfully", Status: "success", Data: map[string]any{"ledger_entries": res}})
}