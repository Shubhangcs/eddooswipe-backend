package repositories

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

type FundInterface interface {
	CreateMasterDistributorFundRequestAdmin(echo.Context) error
	CreateDistributorFundRequestAdmin(echo.Context) error
	CreateRetailerFundRequestAdmin(echo.Context) error
	AcceptFundRequest(echo.Context) error
	RejectFundRequest(echo.Context) error
	GetRequestToFundRequest(c echo.Context) (*[]models.GetFundRequestModel, error)
	GetRequesterFundRequest(c echo.Context) (*[]models.GetFundRequestModel, error)
}

type fundRepository struct {
	db *database.Database
}

func NewFundRepository(db *database.Database) *fundRepository {
	return &fundRepository{
		db,
	}
}

func (fr *fundRepository) CreateMasterDistributorFundRequestAdmin(c echo.Context) error {
	var req models.CreateFundRequestModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.CreateMasterDistributorFundRequestAdminQuery(ctx, req)
}

func (fr *fundRepository) CreateDistributorFundRequestAdmin(c echo.Context) error {
	var req models.CreateFundRequestModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.CreateDistributorFundRequestAdminQuery(ctx, req)
}

func (fr *fundRepository) CreateRetailerFundRequestAdmin(c echo.Context) error {
	var req models.CreateFundRequestModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.CreateRetailerFundRequestAdminQuery(ctx, req)
}

func (fr *fundRepository) AcceptFundRequest(c echo.Context) error {
	var req models.AcceptFundRequestModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.AcceptFundRequestQuery(ctx, req)
}

func (fr *fundRepository) RejectFundRequest(c echo.Context) error {
	var req models.RejectFundRequestModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.RejectFundRequestQuery(ctx, req)
}

func (fr *fundRepository) GetRequestToFundRequest(c echo.Context) (*[]models.GetFundRequestModel, error) {
	var id = c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.GetFundRequestByRequestToID(ctx, id)
}

func (fr *fundRepository) GetRequesterFundRequest(c echo.Context) (*[]models.GetFundRequestModel, error) {
	var id = c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.GetFundRequestByRequesterID(ctx, id)
}
