package models

type CheckMerchantRegistrationResponseModel struct {
	Message      string `json:"message"`
	ResponseCode int    `json:"response_code"`
	Status       bool   `json:"status"`
}

type RegisterMerchantRequest struct {
	MerchantID string `json:"merchantid" validate:"required"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name" validate:"required"`
	PAN        string `json:"pan"`
	Address    string `json:"address"`
	City       string `json:"city"`
	StateCode  string `json:"statecode" validate:"required"`
	PinCode    string `json:"pincode"`
	DOB        string `json:"dob"`
	Email      string `json:"email"`
	FirmName   string `json:"firmname"`
	Latitude   string `json:"latitude" validate:"required"`
	Longitude  string `json:"longitude" validate:"required"`
	Aadhar     string `json:"aadhar"`
	PIDData    string `json:"piddata" validate:"required"`
	AccessMode string `json:"accessmode" validate:"request"`
}
