package repositories

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

type UsersInterface interface {
	GetAllAdmins(echo.Context) (*[]models.GetAdminModel, error)
	GetAllMasterDistributors(echo.Context) (*[]models.GetMasterDistributorModel, error)
	GetAllDistributors(echo.Context) (*[]models.GetDistributorModel, error)
	GetAllRetailers(echo.Context) (*[]models.GetRetailerModel, error)
}

type usersRepository struct {
	db       *database.Database
	jwtUtils *pkg.JWTUtils
}

func NewUsersRepository(db *database.Database, jwtUtils *pkg.JWTUtils) *usersRepository {
	return &usersRepository{
		db,
		jwtUtils,
	}
}

func (ur *usersRepository) GetAllAdmins(c echo.Context) (*[]models.GetAdminModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return ur.db.GetAllAdminsQuery(ctx)
}

func (ur *usersRepository) GetAllMasterDistributors(c echo.Context) (*[]models.GetMasterDistributorModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return ur.db.GetAllMasterDistributorsQuery(ctx)
}

func (ur *usersRepository) GetAllDistributors(c echo.Context) (*[]models.GetDistributorModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return ur.db.GetAllDistributorsQuery(ctx)
}

func (ur *usersRepository) GetAllRetailers(c echo.Context) (*[]models.GetRetailerModel, error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	return ur.db.GetAllRetailersQuery(ctx)
}
