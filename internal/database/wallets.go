package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

func (db *Database) GetAdminWalletBalanceQuery(ctx context.Context, adminID string) (string, error) {
	var adminWalletBalance string
	query := `
		SELECT admin_wallet::TEXT FROM admins WHERE admin_id=@admin_id;
	`
	if err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"admin_id": adminID,
	}).Scan(&adminWalletBalance); err != nil {
		return "", fmt.Errorf("failed to get admin wallet balance")
	}
	return adminWalletBalance, nil
}

func (db *Database) GetMasterDistributorWalletBalanceQuery(ctx context.Context, masterDistributorID string) (string, error) {
	var masterDistributorWalletBalance string
	query := `
		SELECT master_distributor_wallet::TEXT FROM master_distributors WHERE master_distributor_id=@master_distributor_id;
	`
	if err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"master_distributor_id": masterDistributorID,
	}).Scan(&masterDistributorWalletBalance); err != nil {
		return "", fmt.Errorf("failed to get master distributor wallet balance")
	}
	return masterDistributorWalletBalance, nil
}

func (db *Database) GetDistributorWalletBalanceQuery(ctx context.Context, distributorID string) (string, error) {
	var distributorWalletBalance string
	query := `
		SELECT distributor_wallet::TEXT FROM distributors WHERE distributor_id=@distributor_id;
	`
	if err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"distributor_id": distributorID,
	}).Scan(&distributorWalletBalance); err != nil {
		return "", fmt.Errorf("failed to get distributor wallet balance")
	}
	return distributorWalletBalance, nil
}

func (db *Database) GetRetailerWalletBalanceQuery(ctx context.Context, retailerID string) (string, error) {
	var retailerWalletBalance string
	query := `
		SELECT retailer_wallet::TEXT FROM retailers WHERE retailer_id=@retailer_id;
	`
	if err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"retailer_id": retailerID,
	}).Scan(&retailerWalletBalance); err != nil {
		return "", fmt.Errorf("failed to get retailer wallet balance")
	}
	return retailerWalletBalance, nil
}

func (db *Database) AdminWalletTopupQuery(ctx context.Context, req models.AdminWalletTopupModel) error {
	var remarks string
	if req.Remarks == "" {
		remarks = "Admin Wallet Topup"
	}
	remarks = req.Remarks
	adminWalletBalanceUpdateQuery := `
		UPDATE admins
		SET admin_wallet = admin_wallet + @amount::NUMERIC
		WHERE admin_id = @admin_id AND @amount::NUMERIC > 0;
	`
	ledgerEntryQuery := `
		WITH admin_wallet_details AS (
			SELECT admin_wallet FROM admins WHERE admin_id=@admin_id
		)
		INSERT INTO ledger_entries (
			transactor_id,
			reference_id,
			credit_amount,
			latest_balance,
			remarks
		) VALUES (
			@admin_id,
			@ledger_reference_id,
			@ledger_credit_amount,
			(SELECT admin_wallet FROM admin_wallet_details),
			@ledger_remarks
		);
	`
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to topup admin wallet")
	}
	defer tx.Rollback(ctx)

	res, err := tx.Exec(ctx, adminWalletBalanceUpdateQuery, pgx.NamedArgs{
		"admin_id": req.AdminID,
		"amount":   req.Amount,
	})
	if err != nil {
		return fmt.Errorf("failed to update admin wallet balance")
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("failed to topup admin wallet")
	}

	if _, err := tx.Exec(ctx, ledgerEntryQuery, pgx.NamedArgs{
		"admin_id":             req.AdminID,
		"ledger_reference_id":  "NONE",
		"ledger_credit_amount": req.Amount,
		"ledger_remarks":       remarks,
	}); err != nil {
		return fmt.Errorf("failed to topup admin wallet: %w", err)
	}
	return tx.Commit(ctx)
}

func (db *Database) GetLedgerTransactionsQuery(ctx context.Context, id string) (*[]models.GetLedgerEntriesModel, error) {
	query := `
		SELECT ledger_trandaction_id, transactor_id, reference_id, remarks
		credit_amount, debit_amount, latest_balance, created_at
		FROM ledger_entries
		WHERE transactor_id=@id;
	`
	res, err := db.pool.Query(ctx, query, pgx.NamedArgs{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get ledger data")
	}
	defer res.Close()

	var transactions []models.GetLedgerEntriesModel
	for res.Next() {
		var transaction models.GetLedgerEntriesModel
		if err := res.Scan(
			&transaction.TransactionID,
			&transaction.TransactorID,
			&transaction.ReferenceID,
			&transaction.Remarks,
			&transaction.CreditAmount,
			&transaction.DebitAmount,
			&transaction.CreatedAT,
		); err != nil {
			return nil, fmt.Errorf("failed to get ledger data")
		}

		transactions = append(transactions, transaction)
	}

	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get ledger data")
	}
	return &transactions, nil
}
