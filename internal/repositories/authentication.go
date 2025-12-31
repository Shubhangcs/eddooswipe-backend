package repositories

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

type AuthenticationInterface interface {
	CreateAdmin(echo.Context) error
	CreateMasterDistributor(echo.Context) error
	CreateDistributor(echo.Context) error
	CreateRetailer(echo.Context) error
	LoginAdmin(echo.Context) (string, error)
	LoginMasterDistributor(echo.Context) (string, error)
	LoginDistributor(echo.Context) (string, error)
	LoginRetailer(echo.Context) (string, error)
}

type authenticationRepository struct {
	db       *database.Database
	jwtUtils *pkg.JWTUtils
}

func NewAuthenticationRepository(db *database.Database, jwtUtils *pkg.JWTUtils) *authenticationRepository {
	return &authenticationRepository{
		db,
		jwtUtils,
	}
}

func (ar *authenticationRepository) CreateAdmin(c echo.Context) error {
	var req models.CreateAdminModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	if err := ar.db.CreateAdminQuery(ctx, req); err != nil {
		return err
	}
	return nil
}

func (ar *authenticationRepository) CreateMasterDistributor(c echo.Context) error {
	var req models.CreateMasterDistributorModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	if err := ar.db.CreateMasterDistributorQuery(ctx, req); err != nil {
		return err
	}
	return nil
}

func (ar *authenticationRepository) CreateDistributor(c echo.Context) error {
	var req models.CreateDistributorModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	if err := ar.db.CreateDistributorQuery(ctx, req); err != nil {
		return err
	}
	return nil
}

func (ar *authenticationRepository) CreateRetailer(c echo.Context) error {
	var req models.CreateRetailerModel
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	if err := ar.db.CreateRetailerQuery(ctx, req); err != nil {
		return err
	}
	return nil
}

func (ar *authenticationRepository) LoginAdmin(c echo.Context) (string, error) {
	var req models.AdminLoginModel
	if err := bindAndValidate(c, &req); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	res, err := ar.db.LoginAdminQuery(ctx, req)
	if err != nil {
		return "", err
	}
	token, err := ar.jwtUtils.GenerateToken(*res)
	if err != nil {
		return "", err
	}
	log.Println(token)
	return token, nil
}

func (ar *authenticationRepository) LoginMasterDistributor(c echo.Context) (string, error) {
	var req models.MasterDistributorLoginModel
	if err := bindAndValidate(c, &req); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	res, err := ar.db.LoginMasterDistributorQuery(ctx, req)
	if err != nil {
		return "", err
	}
	token, err := ar.jwtUtils.GenerateToken(*res)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ar *authenticationRepository) LoginDistributor(c echo.Context) (string, error) {
	var req models.DistributorLoginModel
	if err := bindAndValidate(c, &req); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	res, err := ar.db.LoginDistributorQuery(ctx, req)
	if err != nil {
		return "", err
	}
	token, err := ar.jwtUtils.GenerateToken(*res)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ar *authenticationRepository) LoginRetailer(c echo.Context) (string, error) {
	var req models.RetailerLoginModel
	if err := bindAndValidate(c, &req); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context() , time.Second * 10)
	defer cancel()
	res, err := ar.db.LoginRetailerQuery(ctx, req)
	if err != nil {
		return "", err
	}
	token, err := ar.jwtUtils.GenerateToken(*res)
	if err != nil {
		return "", err
	}
	return token, nil
}
