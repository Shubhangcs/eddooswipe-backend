package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

func (db *Database) GetRetailerDetailsForDMT(ctx context.Context, req *models.RegisterMerchantRequest) error {
	query := `
		SELECT
			retailer_email,
			retailer_phone,
			retailer_pan_number,
			retailer_aadhar_number,
			retailer_pin_code,
			retailer_city,
			retailer_address,
			retailer_firm_name
		FROM retailers
		WHERE retailer_id=@retailer_id;
	`
	if err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"retailer_id": req.MerchantID,
	}).Scan(
		&req.Email,
		&req.Mobile,
		&req.PAN,
		&req.Aadhar,
		&req.PinCode,
		&req.City,
		&req.Address,
		&req.FirmName,
	); err != nil {
		return fmt.Errorf("failed to fetch retailer details")
	}
	return nil
}
