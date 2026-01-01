package repositories

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

type FundInterface interface {
	CreateFundRequest(echo.Context) error
	AcceptFundRequest(echo.Context) error
	RejectFundRequest(echo.Context) error
	GetAllFundRequests(echo.Context) error
}

type fundRepository struct {
	db *database.Database
}

func NewFundRepository(db *database.Database) *fundRepository {
	return &fundRepository{
		db,
	}
}

func (fr *fundRepository) CreateFundRequest(c echo.Context) error {
	var req models.CreateFundRequest
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.CreateFundRequestQuery(ctx, req)
}

func (fr *fundRepository) AcceptFundRequest(c echo.Context) error {
	var req models.AcceptFundRequest
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return fr.db.AcceptFundRequestQuery(ctx, req)
}


