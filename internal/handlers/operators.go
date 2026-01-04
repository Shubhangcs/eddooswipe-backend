package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type operatorHandler struct {
	operatorRepository repositories.OperatorInterface
}

func NewOperatorHandler(operatorRepository repositories.OperatorInterface) *operatorHandler {
	return &operatorHandler{
		operatorRepository,
	}
}

func (oh *operatorHandler) CreateOperatorRequest(c echo.Context) error {
	err := oh.operatorRepository.CreateOperator(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "operator created successfully", Status: "success"})
}

func (oh *operatorHandler) GetAllOperatorsRequest(c echo.Context) error {
	res, err := oh.operatorRepository.GetAllOperators(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "operators fetched successfully", Status: "success", Data: map[string]any{"operators": res}})
}
