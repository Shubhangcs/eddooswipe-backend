package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

func (db *Database) CreateAdminBankQuery(ctx context.Context, req models.CreateBankModel) error {
	query := `
		INSERT INTO admin_banks(
			admin_id,
			bank_name,
			bank_address,
			bank_account_holder_name,
			bank_account_number,
			bank_ifsc_code
		) VALUES (
			@admin_id,
			@bank_name,
			@bank_address,
			@bank_account_holder_name,
			@bank_account_number,
			@bank_ifsc_code
		);
	`
	if _, err := db.pool.Exec(ctx, query, pgx.NamedArgs{
		"admin_id":                 req.UserID,
		"bank_name":                req.BankName,
		"bank_address":             req.BankAddress,
		"bank_account_holder_name": req.BankAccountHolderName,
		"bank_account_number":      req.BankAccountNumber,
		"bank_ifsc_code":           req.BankIFSCCode,
	}); err != nil {
		return fmt.Errorf("failed to add admin bank")
	}
	return nil
}

func (db *Database) CreateRetailerBankQuery(ctx context.Context, req models.CreateBankModel) error {
	query := `
		INSERT INTO retailer_banks(
			retailer_id,
			bank_name,
			bank_address,
			bank_account_holder_name,
			bank_account_number,
			bank_ifsc_code
		) VALUES (
			@retailer_id,
			@bank_name,
			@bank_address,
			@bank_account_holder_name,
			@bank_account_number,
			@bank_ifsc_code
		);
	`
	if _, err := db.pool.Exec(ctx, query, pgx.NamedArgs{
		"retailer_id":              req.UserID,
		"bank_name":                req.BankName,
		"bank_address":             req.BankAddress,
		"bank_account_holder_name": req.BankAccountHolderName,
		"bank_account_number":      req.BankAccountNumber,
		"bank_ifsc_code":           req.BankIFSCCode,
	}); err != nil {
		return fmt.Errorf("failed to add retailer bank")
	}
	return nil
}

func (db *Database) GetAdminBanksByAdminIDQuery(ctx context.Context, adminID string) (*[]models.GetBanksModel, error) {
	query := `
		SELECT admin_id, bank_name, bank_address, bank_account_holder_name,
		bank_account_number, bank_ifsc_code, created_at, updated_at
		FROM admin_banks
		WHERE admin_id=@admin_id;
	`
	res, err := db.pool.Query(ctx, query, pgx.NamedArgs{
		"admin_id": adminID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get admin bank data")
	}
	defer res.Close()

	var adminBanks []models.GetBanksModel
	for res.Next() {
		var adminBank models.GetBanksModel
		if err := res.Scan(
			&adminBank.UserID,
			&adminBank.BankName,
			&adminBank.BankAddress,
			&adminBank.BankAccountHolderName,
			&adminBank.BankAccountNumber,
			&adminBank.BankIFSCCode,
			&adminBank.CreatedAt,
			&adminBank.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get admin bank data")
		}
		adminBanks = append(adminBanks, adminBank)
	}
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get admin bank data")
	}
	return &adminBanks, nil
}

func (db *Database) GetRetailerBanksByRetailerIDQuery(ctx context.Context, retailerID string) (*[]models.GetBanksModel, error) {
	query := `
		SELECT retailer_id, bank_name, bank_address, bank_account_holder_name,
		bank_account_number, bank_ifsc_code, created_at, updated_at
		FROM retailer_banks
		WHERE retailer_id=@retailer_id;
	`
	res, err := db.pool.Query(ctx, query, pgx.NamedArgs{
		"retailer_id": retailerID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get retailer bank data")
	}
	defer res.Close()

	var retailerBanks []models.GetBanksModel
	for res.Next() {
		var retailerBank models.GetBanksModel
		if err := res.Scan(
			&retailerBank.UserID,
			&retailerBank.BankName,
			&retailerBank.BankAddress,
			&retailerBank.BankAccountHolderName,
			&retailerBank.BankAccountNumber,
			&retailerBank.BankIFSCCode,
			&retailerBank.CreatedAt,
			&retailerBank.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get retailer bank data")
		}
		retailerBanks = append(retailerBanks, retailerBank)
	}
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get retailer bank data")
	}
	return &retailerBanks, nil
}

func (db *Database) GetAllAdminBanksQuery(ctx context.Context) (*[]models.GetBanksModel, error) {
	query := `
		SELECT admin_id, bank_name, bank_address, bank_account_holder_name,
		bank_account_number, bank_ifsc_code, created_at, updated_at
		FROM admin_banks;
	`
	res, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin bank data")
	}
	defer res.Close()

	var adminBanks []models.GetBanksModel
	for res.Next() {
		var adminBank models.GetBanksModel
		if err := res.Scan(
			&adminBank.UserID,
			&adminBank.BankName,
			&adminBank.BankAddress,
			&adminBank.BankAccountHolderName,
			&adminBank.BankAccountNumber,
			&adminBank.BankIFSCCode,
			&adminBank.CreatedAt,
			&adminBank.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get admin bank data")
		}
		adminBanks = append(adminBanks, adminBank)
	}
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get admin bank data")
	}
	return &adminBanks, nil
}

func (db *Database) GetAllRetailerBanksQuery(ctx context.Context) (*[]models.GetBanksModel, error) {
	query := `
		SELECT retailer_id, bank_name, bank_address, bank_account_holder_name,
		bank_account_number, bank_ifsc_code, created_at, updated_at
		FROM retailer_banks;
	`
	res, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get retailer bank data")
	}
	defer res.Close()

	var retailerBanks []models.GetBanksModel
	for res.Next() {
		var retailerBank models.GetBanksModel
		if err := res.Scan(
			&retailerBank.UserID,
			&retailerBank.BankName,
			&retailerBank.BankAddress,
			&retailerBank.BankAccountHolderName,
			&retailerBank.BankAccountNumber,
			&retailerBank.BankIFSCCode,
			&retailerBank.CreatedAt,
			&retailerBank.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get retailer bank data")
		}
		retailerBanks = append(retailerBanks, retailerBank)
	}
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get retailer bank data")
	}
	return &retailerBanks, nil
}
