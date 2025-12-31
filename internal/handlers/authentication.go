package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type authenticationHandler struct {
	authenticationRepository repositories.AuthenticationInterface
}

func NewAuthenticationHandler(authenticationRepository repositories.AuthenticationInterface) *authenticationHandler {
	return &authenticationHandler{
		authenticationRepository,
	}
}

func (ah *authenticationHandler) CreateAdminRequest(c echo.Context) error {
	if err := ah.authenticationRepository.CreateAdmin(c); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "admin created successfully", Status: "success"})
}

func (ah *authenticationHandler) CreateMasterDistributorRequest(c echo.Context) error {
	if err := ah.authenticationRepository.CreateMasterDistributor(c); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "master distributor created successfully", Status: "success"})
}

func (ah *authenticationHandler) CreateDistributorRequest(c echo.Context) error {
	if err := ah.authenticationRepository.CreateDistributor(c); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "distributor created successfully", Status: "success"})
}

func (ah *authenticationHandler) CreateRetailerRequest(c echo.Context) error {
	if err := ah.authenticationRepository.CreateRetailer(c); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "retailer created successfully", Status: "success"})
}

func (ah *authenticationHandler) LoginAdminRequest(c echo.Context) error {
	res, err := ah.authenticationRepository.LoginAdmin(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "admin login successfull", Status: "success", Data: map[string]string{"token": res}})
}

func (ah *authenticationHandler) LoginMasterDistributorRequest(c echo.Context) error {
	res, err := ah.authenticationRepository.LoginMasterDistributor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "master distributor login successfull", Status: "success", Data: map[string]string{"token": res}})
}

func (ah *authenticationHandler) LoginDistributorRequest(c echo.Context) error {
	res, err := ah.authenticationRepository.LoginDistributor(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "distributor login successfull", Status: "success", Data: map[string]string{"token": res}})
}

func (ah *authenticationHandler) LoginRetailerRequest(c echo.Context) error {
	res, err := ah.authenticationRepository.LoginRetailer(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusCreated, models.ResponseModel{Message: "retailer login successfull", Status: "success", Data: map[string]string{"token": res}})
}
