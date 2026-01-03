package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

func (db *Database) CreateFundRequestQuery(ctx context.Context, req models.CreateFundRequestModel) error {
	query := `
		INSERT INTO fund_requests(
			requester_id,
			requester_name,
			request_to_id,
			request_to_name,
			amount,
			account_number,
			bank_name,
			deposit_date,
			utr_number,
			payment_mode,
			fund_request_status,
			requester_remarks
		) VALUES (
			@requester_id,
			@requester_name,
			@request_to_id,
			@request_to_name,
			@amount,
			@account_number,
			@bank_name,
			@deposit_date,
			@utr_number,
			@payment_mode,
			'PENDING',
			@requester_remarks 
		);
	`
	if _, err := db.pool.Exec(ctx, query, pgx.NamedArgs{
		"requester_id":      req.RequesterID,
		"requester_name":    req.RequesterName,
		"request_to_id":     req.RequestToID,
		"request_to_name":   req.RequestToName,
		"amount":            req.Amount,
		"account_number":    req.AccountNumber,
		"bank_name":         req.BankName,
		"deposit_date":      req.DepositDate,
		"utr_number":        req.UTRNumber,
		"payment_mode":      req.PaymentMode,
		"requester_remarks": req.RequesterRemarks,
	}); err != nil {
		return fmt.Errorf("failed to create fund request: %w", err)
	}
	return nil
}

func getLatestBalanceTx(ctx context.Context, tx pgx.Tx, table, walletCol, idCol, userID string) (string, error) {

	query := fmt.Sprintf(
		"SELECT %s::TEXT FROM %s WHERE %s = @user_id",
		walletCol,
		table,
		idCol,
	)

	var balance string
	err := tx.QueryRow(
		ctx,
		query,
		pgx.NamedArgs{"user_id": userID},
	).Scan(&balance)

	return balance, err
}

func (db *Database) AcceptFundRequestQuery(ctx context.Context, req models.AcceptFundRequestModel) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start fund request transaction")
	}
	defer tx.Rollback(ctx)

	var fundRequestDetails struct {
		RequesterID string `json:"requester_id"`
		RequestToID string `json:"request_to_id"`
		Amount      string `json:"amount"`
	}
	getFundRequestDetails := `
		SELECT requester_id, request_to_id, amount FROM fund_requests WHERE fund_request_id=@fund_request_id;
	`

	if err := tx.QueryRow(ctx, getFundRequestDetails, pgx.NamedArgs{
		"fund_request_id": req.FundRequestID,
	}).Scan(
		&fundRequestDetails.RequesterID,
		&fundRequestDetails.RequestToID,
		&fundRequestDetails.Amount,
	); err != nil {
		return fmt.Errorf("failed to accept fund request")
	}

	type WalletTableDetails struct {
		TableName        string
		TableIDFieldName string
		TableWalletName  string
	}

	walletTableMap := map[byte]WalletTableDetails{
		'A': {"admins", "admin_id", "admin_wallet"},
		'M': {"master_distributors", "master_distributor_id", "master_distributor_wallet"},
		'D': {"distributors", "distributor_id", "distributor_wallet"},
		'R': {"retailers", "retailer_id", "retailer_wallet"},
	}

	requestTo, ok := walletTableMap[fundRequestDetails.RequestToID[0]]
	if !ok {
		return fmt.Errorf("invalid request_to id")
	}

	requester, ok := walletTableMap[fundRequestDetails.RequesterID[0]]
	if !ok {
		return fmt.Errorf("invalid requester id")
	}

	/* -------------------- DEBIT -------------------- */

	deductQuery := fmt.Sprintf(`
		UPDATE %s
		SET %s = %s - @amount
		WHERE %s = @user_id AND %s >= @amount
	`,
		requestTo.TableName,
		requestTo.TableWalletName,
		requestTo.TableWalletName,
		requestTo.TableIDFieldName,
		requestTo.TableWalletName,
	)

	tag, err := tx.Exec(ctx, deductQuery, pgx.NamedArgs{
		"user_id": fundRequestDetails.RequestToID,
		"amount":  fundRequestDetails.Amount,
	})
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("insufficient balance")
	}

	debitBalance, err := getLatestBalanceTx(
		ctx, tx,
		requestTo.TableName,
		requestTo.TableWalletName,
		requestTo.TableIDFieldName,
		fundRequestDetails.RequestToID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO ledger_entries (
			transactor_id,
			reference_id,
			debit_amount,
			remarks,
			latest_balance
		) VALUES (
			@transactor_id,
			@reference_id,
			@debit_amount,
			@remarks,
			@latest_balance
		)
	`, pgx.NamedArgs{
		"transactor_id":  fundRequestDetails.RequestToID,
		"reference_id":   req.FundRequestID,
		"debit_amount":   fundRequestDetails.Amount,
		"remarks":        fmt.Sprintf("Fund transferred to %s", fundRequestDetails.RequesterID),
		"latest_balance": debitBalance,
	})
	if err != nil {
		return err
	}

	/* -------------------- CREDIT -------------------- */

	creditQuery := fmt.Sprintf(`
		UPDATE %s
		SET %s = %s + @amount
		WHERE %s = @user_id
	`,
		requester.TableName,
		requester.TableWalletName,
		requester.TableWalletName,
		requester.TableIDFieldName,
	)

	_, err = tx.Exec(ctx, creditQuery, pgx.NamedArgs{
		"user_id": fundRequestDetails.RequesterID,
		"amount":  fundRequestDetails.Amount,
	})
	if err != nil {
		return err
	}

	creditBalance, err := getLatestBalanceTx(
		ctx, tx,
		requester.TableName,
		requester.TableWalletName,
		requester.TableIDFieldName,
		fundRequestDetails.RequesterID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO ledger_entries (
			transactor_id,
			reference_id,
			credit_amount,
			remarks,
			latest_balance
		) VALUES (
			@transactor_id,
			@reference_id,
			@credit_amount,
			@remarks,
			@latest_balance
		)
	`, pgx.NamedArgs{
		"transactor_id":  fundRequestDetails.RequesterID,
		"reference_id":   req.FundRequestID,
		"credit_amount":  fundRequestDetails.Amount,
		"remarks":        fmt.Sprintf("Fund received from %s", fundRequestDetails.RequestToID),
		"latest_balance": creditBalance,
	})
	if err != nil {
		return err
	}

	/* -------------------- FUND REQUEST STATUS -------------------- */

	_, err = tx.Exec(ctx, `
		UPDATE fund_requests
		SET fund_request_status = 'ACCEPTED',
		    request_to_remarks = @remarks
		WHERE fund_request_id = @id
	`, pgx.NamedArgs{
		"id":      req.FundRequestID,
		"remarks": req.RequestToRemarks,
	})
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (db *Database) RejectFundRequestQuery(ctx context.Context, req models.RejectFundRequestModel) error {
	query := `
		UPDATE fund_requests 
		SET fund_request_status = 'REJECTED' , request_to_remarks=@request_to_remarks 
		WHERE fund_request_id=@fund_request_id;
	`
	if _, err := db.pool.Exec(ctx, query, pgx.NamedArgs{
		"fund_request_id":    req.FundRequestID,
		"request_to_remarks": req.RequestToRemarks,
	}); err != nil {
		return fmt.Errorf("failed to reject fund request")
	}
	return nil
}

func (db *Database) GetFundRequestByRequestToID(ctx context.Context, requestToID string) (*[]models.GetFundRequestModel, error) {
	query := `
		SELECT 
			request_id,
			requester_id,
			requester_name,
			request_to_id,
			request_to_name,
			account_number,
			bank_name,
			amount,
			fund_request_status,
			payment_mode,
			utr_number,
			deposit_date,
			requester_remarks,
			request_to_remarks,
			created_at,
			updated_at
		FROM fund_requests
		WHERE request_to_id=@request_to_id;
	`
	res, err := db.pool.Query(ctx, query, pgx.NamedArgs{
		"request_to_id": requestToID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get fund request data")
	}
	defer res.Close()

	var fundRequests []models.GetFundRequestModel
	for res.Next() {
		var fundRequest models.GetFundRequestModel
		if err := res.Scan(
			&fundRequest.RequestID,
			&fundRequest.RequesterID,
			&fundRequest.RequesterName,
			&fundRequest.RequestToID,
			&fundRequest.RequestToName,
			&fundRequest.AccountNumber,
			&fundRequest.BankName,
			&fundRequest.Amount,
			&fundRequest.FundRequestStatus,
			&fundRequest.PaymentMode,
			&fundRequest.UTRNumber,
			&fundRequest.DepositDate,
			&fundRequest.DepositDate,
			&fundRequest.RequesterRemarks,
			&fundRequest.RequestToRemarks,
			&fundRequest.CreatedAt,
			&fundRequest.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get fund request data")
		}
		fundRequests = append(fundRequests, fundRequest)
	}

	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get fund request data")
	}
	return &fundRequests, nil
}

func (db *Database) GetFundRequestByRequesterID(ctx context.Context, requesterID string) (*[]models.GetFundRequestModel, error) {
	query := `
		SELECT 
			request_id,
			requester_id,
			requester_name,
			request_to_id,
			request_to_name,
			account_number,
			bank_name,
			amount,
			fund_request_status,
			payment_mode,
			utr_number,
			deposit_date,
			requester_remarks,
			request_to_remarks,
			created_at,
			updated_at
		FROM fund_requests
		WHERE requester_id=@requester_id;
	`
	res, err := db.pool.Query(ctx, query, pgx.NamedArgs{
		"requester_id": requesterID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get fund request data")
	}
	defer res.Close()

	var fundRequests []models.GetFundRequestModel
	for res.Next() {
		var fundRequest models.GetFundRequestModel
		if err := res.Scan(
			&fundRequest.RequestID,
			&fundRequest.RequesterID,
			&fundRequest.RequesterName,
			&fundRequest.RequestToID,
			&fundRequest.RequestToName,
			&fundRequest.AccountNumber,
			&fundRequest.BankName,
			&fundRequest.Amount,
			&fundRequest.FundRequestStatus,
			&fundRequest.PaymentMode,
			&fundRequest.UTRNumber,
			&fundRequest.DepositDate,
			&fundRequest.DepositDate,
			&fundRequest.RequesterRemarks,
			&fundRequest.RequestToRemarks,
			&fundRequest.CreatedAt,
			&fundRequest.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get fund request data")
		}
		fundRequests = append(fundRequests, fundRequest)
	}

	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get fund request data")
	}
	return &fundRequests, nil
}
