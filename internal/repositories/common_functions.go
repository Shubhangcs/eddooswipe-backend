package repositories

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func bindAndValidate(c echo.Context, req any) error {
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("invalid request format")
	}
	if err := c.Validate(req); err != nil {
		return fmt.Errorf("invalid request data")
	}
	return nil
}