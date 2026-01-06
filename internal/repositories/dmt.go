package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

type DMTInterface interface {
	CheckMerchantRegistration(c echo.Context) (*models.CheckMerchantRegistrationResponseModel, error)
	RegisterMerchant(echo.Context) (any, error)
}

type dmtRepository struct {
	db           *database.Database
	jwtUtils     *pkg.JWTUtils
	paysprintURL string
}

func NewDMTRepository(db *database.Database, jwtUtils *pkg.JWTUtils, paysprintURL string) *dmtRepository {
	return &dmtRepository{
		db,
		jwtUtils,
		paysprintURL,
	}
}

func (dr *dmtRepository) CheckMerchantRegistration(c echo.Context) (*models.CheckMerchantRegistrationResponseModel, error) {
	var retailerID = c.Param("retailer_id")
	jsonBytes, err := json.Marshal(map[string]string{"merchantid": retailerID})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json")
	}
	payload := bytes.NewReader(jsonBytes)
	var url = fmt.Sprintf("%s/service/dmt-v6/merchant/status_check", dr.paysprintURL)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request")
	}

	token, err := dr.jwtUtils.GenerateTokenForPaysprint()
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorized_key", "UFMwMDI3NDZmZjUyNjIzZmM3OGM2MzJhYWIwMTAzYmRjZjFlYTgzMQ==")
	req.Header.Add("Token", token)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response models.CheckMerchantRegistrationResponseModel
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, err
}

func (dr *dmtRepository) RegisterMerchant(c echo.Context) (any, error) {
	var req models.RegisterMerchantRequest
	if err := bindAndValidate(c, &req); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*10)
	defer cancel()
	if err := dr.db.GetRetailerDetailsForDMT(ctx, &req); err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json")
	}
	payload := bytes.NewReader(jsonBytes)
	var url = fmt.Sprintf("%s/service/dmt-v6/merchant/register", dr.paysprintURL)

	r, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to process request")
	}
	token, err := dr.jwtUtils.GenerateTokenForPaysprint()
	if err != nil {
		return nil, err
	}
	r.Header.Add("accept", "application/json")
	r.Header.Add("authorized_key", "UFMwMDI3NDZmZjUyNjIzZmM3OGM2MzJhYWIwMTAzYmRjZjFlYTgzMQ==")
	r.Header.Add("Token", token)
	r.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response any
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
