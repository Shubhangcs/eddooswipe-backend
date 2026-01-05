package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/internal/repositories"
)

type dmtHandler struct {
	dmtRepository repositories.DMTInterface
}

func NewDMTHandler(dmtRepository repositories.DMTInterface) *dmtHandler {
	return &dmtHandler{
		dmtRepository,
	}
}

func (dh *dmtHandler) RegisterMerchantRequest(c echo.Context) error {
	res, err := dh.dmtRepository.RegisterMerchant(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseModel{Message: err.Error(), Status: "failed"})
	}
	return c.JSON(http.StatusOK, models.ResponseModel{Message: "dmt register merchant check success", Status: "success", Data: map[string]any{"response": res}})
}
