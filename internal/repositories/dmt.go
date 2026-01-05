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
	url := "https://api.paysprint.in/api/v1/service/dmt/kyc/remitter/queryremitter"

	payload := strings.NewReader("{\"mobile\":9773870841}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorized_key", "UFMwMDI3NDZmZjUyNjIzZmM3OGM2MzJhYWIwMTAzYmRjZjFlYTgzMQ==")
	req.Header.Add("Token" , "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXJ0bmVySWQiOiJQUzAwMjc0NiIsInJlcWlkIjoiNTY1NDQ1NjciLCJ0aW1lc3RhbXAiOjE3Njc2MzU2NTB9.ohQSa6YTI97yJF5C7fJxG7K0iGa8kR925RMWmAEDMFk")
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
