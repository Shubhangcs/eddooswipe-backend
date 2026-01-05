package repositories

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
)

type DMTInterface interface {
	RegisterMerchant(echo.Context) (any, error)
}

type dmtRepository struct {
	db *database.Database
}

func NewDMTRepository(db *database.Database) *dmtRepository {
	return &dmtRepository{
		db,
	}
}

func (dr *dmtRepository) RegisterMerchant(c echo.Context) (any, error) {
	url := "https://sit.paysprint.in/service-api/api/v1/service/dmt/kyc/remitter/queryremitter"

	payload := strings.NewReader("{\"mobile\":9773870841}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJQQVlTUFJJTlQiLCJ0aW1lc3RhbXAiOjE2MTAwMjYzMzgsInBhcnRuZXJJZCI6IlBTMDAxIiwicHJvZHVjdCI6IldBTExFVCIsInJlcWlkIjoxNjEwMDI2MzM4fQ.buzD40O8X_41RmJ0PCYbBYx3IBlsmNb9iVmrVH9Ix64")
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
