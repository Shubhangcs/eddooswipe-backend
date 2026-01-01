package repositories

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

type WalletInterface interface {
	GetAdminWalletBalance(echo.Context) (string, error)
	GetMasterDistributorWalletBalance(echo.Context) (string, error)
	GetDistributorWalletBalance(echo.Context) (string, error)
	GetRetailerWalletBalance(echo.Context) (string, error)
	AdminWalletTopup(echo.Context) error
}

type walletRepository struct {
	db *database.Database
}

func NewWalletRepository(db *database.Database) *walletRepository {
	return &walletRepository{
		db,
	}
}

func (wr *walletRepository) GetAdminWalletBalance(c echo.Context) (string, error) {
	adminID := c.Param("admin_id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return wr.db.GetAdminWalletBalanceQuery(ctx, adminID)
}

func (wr *walletRepository) GetMasterDistributorWalletBalance(c echo.Context) (string, error) {
	masterDistributorID := c.Param("master_distributor_id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return wr.db.GetMasterDistributorWalletBalanceQuery(ctx, masterDistributorID)
}

func (wr *walletRepository) GetDistributorWalletBalance(c echo.Context) (string, error) {
	distributorID := c.Param("distributor_id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return wr.db.GetDistributorWalletBalanceQuery(ctx, distributorID)
}

func (wr *walletRepository) GetRetailerWalletBalance(c echo.Context) (string, error) {
	retailerID := c.Param("retailer_id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return wr.db.GetRetailerWalletBalanceQuery(ctx, retailerID)
}

func (wr *walletRepository) AdminWalletTopup(c echo.Context) error {
	var req models.AdminWalletTopupModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return wr.db.AdminWalletTopupQuery(ctx, req)
}
