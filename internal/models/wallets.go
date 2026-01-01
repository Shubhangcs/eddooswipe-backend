package models

type AdminWalletTopupModel struct {
	AdminID string `json:"admin_id" validate:"required"`
	Amount  string `json:"amount" validate:"required"`
	Remarks string `json:"remarks"`
}
