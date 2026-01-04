package repositories

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

type CommisionInterface interface {
	CreateCommision(echo.Context) error
	GetAllCommisions(echo.Context) (*[]models.GetCommisionsModel, error)
	UpdateCommision(echo.Context) error
	DeleteCommision(echo.Context) error
}

type commisionRepository struct {
	db *database.Database
}

func NewCommisionRepository(db *database.Database) *commisionRepository {
	return &commisionRepository{
		db,
	}
}

func (cr *commisionRepository) CreateCommision(c echo.Context) error {
	var req models.CreateCommisionModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return cr.db.CreateCommisionQuery(ctx, req)
}

func (cr *commisionRepository) GetAllCommisions(c echo.Context) (*[]models.GetCommisionsModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return cr.db.GetAllCommisionsQuery(ctx)
}

func (cr *commisionRepository) UpdateCommision(c echo.Context) error {
	var req models.UpdateCommisionModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return cr.db.UpdateCommisionQuery(ctx, req)
}

func (cr *commisionRepository) DeleteCommision(c echo.Context) error {
	var commisionID = c.Param("commision_id")
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return cr.db.DeleteCommisionQuery(ctx, commisionID)
}
