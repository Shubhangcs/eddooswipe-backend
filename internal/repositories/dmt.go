package repositories

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

type DMTInterface interface {
	RegisterMerchant(echo.Context) (any, error)
}

type dmtRepository struct {
	db       *database.Database
	jwtUtils *pkg.JWTUtils
}

func NewDMTRepository(db *database.Database, jwtUtils *pkg.JWTUtils) *dmtRepository {
	return &dmtRepository{
		db,
		jwtUtils,
	}
}

func (dr *dmtRepository) RegisterMerchant(c echo.Context) (any, error) {
	url := "https://api.paysprint.in/api/v1/service/dmt-v6/merchant/status_check"

	payload := strings.NewReader("{\"merchantcode\":PS002746}")

	req, _ := http.NewRequest("POST", url, payload)
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

	var response any
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
