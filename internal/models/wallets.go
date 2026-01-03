package models

type AdminWalletTopupModel struct {
	AdminID string `json:"admin_id" validate:"required"`
	Amount  string `json:"amount" validate:"required"`
	Remarks string `json:"remarks"`
}

type GetLedgerEntriesModel struct {
	TransactionID string `json:"transaction_id"`
	ReferenceID   string `json:"reference_id"`
	TransactorID  string `json:"transactor_id"`
	Remarks       string `json:"remarks"`
	CreditAmount  string `json:"credit_amount"`
	DebitAmount   string `json:"debit_amount"`
	CreatedAT     string `json:"created_at"`
}
