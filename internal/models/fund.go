package models

type CreateFundRequest struct {
	RequesterID      string `json:"requester_id" validate:"required"`
	RequestToID      string `json:"request_to_id" validate:"required"`
	PaymentMode      string `json:"payment_mode" validate:"required"`
	DepositDate      string `json:"deposit_date" validate:"required"`
	Amount           string `json:"amount" validate:"required"`
	AccountNumber    string `json:"account_number" validate:"required"`
	UTRNumber        string `json:"utr_number" validate:"required"`
	RequesterRemarks string `json:"remarks"`
}

type AcceptFundRequest struct {
	FundRequestID    string `json:"fund_request_id" validate:"required"`
	RequestToRemarks string `json:"request_to_remarks"`
}

type RejectFundRequest struct {
	FundRequestID    string `json:"fund_request_id" validate:"required"`
	RequestToRemarks string `json:"request_to_remarks" validate:"required"`
}

type GetFundRequests struct {
	RequestID         string `json:"request_id"`
	RequesterID       string `json:"requester_id"`
	RequesterName     string `json:"requester_name"`
	RequestToID       string `json:"request_to_id"`
	RequestToName     string `json:"request_to_name"`
	PaymentMode       string `json:"payment_mode"`
	Amount            string `json:"amount"`
	FundRequestStatus string `json:"fund_request_status"`
	BankName          string `json:"bank_name"`
	AccountNumber     string `json:"account_number"`
	UTRNumber         string `json:"utr_number"`
	DepositDate       string `json:"deposit_date"`
	RequesterRemarks  string `json:"requester_remarks"`
	RequestToRemarks  string `json:"request_to_remarks"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
