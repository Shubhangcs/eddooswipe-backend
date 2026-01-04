package repositories

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

type OperatorInterface interface {
	CreateOperator(echo.Context) error
	GetAllOperators(echo.Context) (*[]models.GetOperatorsModel, error)
}

type operatorRepository struct {
	db *database.Database
}

func NewOperatorRepository(db *database.Database) *operatorRepository {
	return &operatorRepository{
		db,
	}
}

func (or *operatorRepository) CreateOperator(c echo.Context) error {
	var req models.CreateOperatorModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return or.db.CreateOperatorQuery(ctx, req.OperatorName)
}

func (or *operatorRepository) GetAllOperators(c echo.Context) (*[]models.GetOperatorsModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return or.db.GetAllOperators(ctx)
}
