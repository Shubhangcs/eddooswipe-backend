package repositories

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

type BankInterface interface {
	CreateAdminBank(echo.Context) error
	CreateRetailerBank(echo.Context) error
	GetAdminBanksByAdminID(echo.Context) (*[]models.GetBanksModel, error)
	GetRetailerBanksByRetailerID(echo.Context) (*[]models.GetBanksModel, error)
	GetAllAdminBanks(echo.Context) (*[]models.GetBanksModel, error)
	GetAllRetailerBanks(echo.Context) (*[]models.GetBanksModel, error)
	DeleteAdminBank(echo.Context) error
	DeleteRetailerBank(echo.Context) error
}

type bankRepository struct {
	db *database.Database
}

func NewBankRepository(db *database.Database) *bankRepository {
	return &bankRepository{
		db,
	}
}

func (br *bankRepository) CreateAdminBank(c echo.Context) error {
	var req models.CreateBankModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.CreateAdminBankQuery(ctx, req)
}

func (br *bankRepository) CreateRetailerBank(c echo.Context) error {
	var req models.CreateBankModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.CreateRetailerBankQuery(ctx, req)
}

func (br *bankRepository) GetAdminBanksByAdminID(c echo.Context) (*[]models.GetBanksModel, error) {
	var adminID = c.Param("admin_id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.GetAdminBanksByAdminIDQuery(ctx, adminID)
}

func (br *bankRepository) GetRetailerBanksByRetailerID(c echo.Context) (*[]models.GetBanksModel, error) {
	var retaileID = c.Param("retailer_id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.GetRetailerBanksByRetailerIDQuery(ctx, retaileID)
}

func (br *bankRepository) GetAllAdminBanks(c echo.Context) (*[]models.GetBanksModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.GetAllAdminBanksQuery(ctx)
}

func (br *bankRepository) GetAllRetailerBanks(c echo.Context) (*[]models.GetBanksModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.GetAllRetailerBanksQuery(ctx)
}

func (br *bankRepository) DeleteAdminBank(c echo.Context) error {
	var accountNumber = c.Param("account_number")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.DeleteAdminBankQuery(ctx, accountNumber)
}

func (br *bankRepository) DeleteRetailerBank(c echo.Context) error {
	var accountNumber = c.Param("account_number")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return br.db.DeleteRetailerBankQuery(ctx, accountNumber)
}
