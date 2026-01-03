package models

type CreateBankModel struct {
	UserID                string `json:"user_id" validate:"required"`
	BankName              string `json:"bank_name" validate:"required"`
	BankAddress           string `json:"bank_address" validate:"required"`
	BankAccountHolderName string `json:"bank_account_holder_name" validate:"required"`
	BankAccountNumber     string `json:"bank_account_number" validate:"required"`
	BankIFSCCode          string `json:"bank_ifsc_code" validate:"required"`
}

type GetBanksModel struct {
	UserID                string `json:"user_id"`
	BankName              string `json:"bank_name"`
	BankAddress           string `json:"bank_address"`
	BankAccountHolderName string `json:"bank_account_holder_name"`
	BankAccountNumber     string `json:"bank_account_number"`
	BankIFSCCode          string `json:"bank_ifsc_code"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}
