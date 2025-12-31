package models

type CreateAdminModel struct {
	AdminName     string `json:"admin_name" validate:"required"`
	AdminEmail    string `json:"admin_email" validate:"required,email"`
	AdminPhone    string `json:"admin_phone" validate:"required,phone"`
	AdminPassword string `json:"admin_password" validate:"required,strpwd"`
}

type CreateMasterDistributorModel struct {
	AdminID                             string `json:"admin_id" validate:"required"`
	MasterDistributorName               string `json:"master_distributor_name" validate:"required"`
	MasterDistributorFatherOrSpouseName string `json:"master_distributor_father_or_spouse_name,omitempty"`
	MasterDistributorEmail              string `json:"master_distributor_email" validate:"required,email"`
	MasterDistributorPhone              string `json:"master_distributor_phone" validate:"required,phone"`
	MasterDistributorPassword           string `json:"master_distributor_password" validate:"required,strpwd"`
	MasterDistributorPanNumber          string `json:"master_distributor_pan_number" validate:"required"`
	MasterDistributorAadharNumber       string `json:"master_distributor_aadhar_number" validate:"required"`
	MasterDistributorGSTNumber          string `json:"master_distributor_gst_number,omitempty"`
	MasterDistributorPinCode            string `json:"master_distributor_pin_code" validate:"required"`
	MasterDistributorCity               string `json:"master_distributor_city" validate:"required"`
	MasterDistributorState              string `json:"master_distributor_state" validate:"required"`
	MasterDistributorAddress            string `json:"master_distributor_address" validate:"required"`
	MasterDistributorFirmName           string `json:"master_distributor_firm_name" validate:"required"`
	MasterDistributorFirmAddress        string `json:"master_distributor_firm_address" validate:"required"`
	MasterDistributorFirmCity           string `json:"master_distributor_firm_city" validate:"required"`
	MasterDistributorFirmPin            string `json:"master_distributor_firm_pin" validate:"required"`
	MasterDistributorFirmState          string `json:"master_distributor_firm_state" validate:"required"`
	MasterDistributorFirmDistrict       string `json:"master_distributor_firm_district" validate:"required"`
}

type CreateDistributorModel struct {
	MasterDistributorID           string `json:"master_distributor_id" validate:"required"`
	DistributorName               string `json:"distributor_name" validate:"required"`
	DistributorFatherOrSpouseName string `json:"distributor_father_or_spouse_name,omitempty"`
	DistributorEmail              string `json:"distributor_email" validate:"required,email"`
	DistributorPhone              string `json:"distributor_phone" validate:"required,phone"`
	DistributorPassword           string `json:"distributor_password" validate:"required,strpwd"`
	DistributorPanNumber          string `json:"distributor_pan_number" validate:"required"`
	DistributorAadharNumber       string `json:"distributor_aadhar_number" validate:"required"`
	DistributorGSTNumber          string `json:"distributor_gst_number,omitempty"`
	DistributorPinCode            string `json:"distributor_pin_code" validate:"required"`
	DistributorCity               string `json:"distributor_city" validate:"required"`
	DistributorState              string `json:"distributor_state" validate:"required"`
	DistributorAddress            string `json:"distributor_address" validate:"required"`
	DistributorFirmName           string `json:"distributor_firm_name" validate:"required"`
	DistributorFirmAddress        string `json:"distributor_firm_address" validate:"required"`
	DistributorFirmCity           string `json:"distributor_firm_city" validate:"required"`
	DistributorFirmPin            string `json:"distributor_firm_pin" validate:"required"`
	DistributorFirmState          string `json:"distributor_firm_state" validate:"required"`
	DistributorFirmDistrict       string `json:"distributor_firm_district" validate:"required"`
}

type CreateRetailerModel struct {
	DistributorID              string `json:"distributor_id" validate:"required"`
	RetailerName               string `json:"retailer_name" validate:"required"`
	RetailerFatherOrSpouseName string `json:"retailer_father_or_spouse_name,omitempty"`
	RetailerEmail              string `json:"retailer_email" validate:"required,email"`
	RetailerPhone              string `json:"retailer_phone" validate:"required,phone"`
	RetailerPassword           string `json:"retailer_password" validate:"required,strpwd"`
	RetailerPanNumber          string `json:"retailer_pan_number" validate:"required"`
	RetailerAadharNumber       string `json:"retailer_aadhar_number" validate:"required"`
	RetailerGSTNumber          string `json:"retailer_gst_number,omitempty"`
	RetailerPinCode            string `json:"retailer_pin_code" validate:"required"`
	RetailerCity               string `json:"retailer_city" validate:"required"`
	RetailerState              string `json:"retailer_state" validate:"required"`
	RetailerAddress            string `json:"retailer_address" validate:"required"`
	RetailerFirmName           string `json:"retailer_firm_name" validate:"required"`
	RetailerFirmAddress        string `json:"retailer_firm_address" validate:"required"`
	RetailerFirmCity           string `json:"retailer_firm_city" validate:"required"`
	RetailerFirmPin            string `json:"retailer_firm_pin" validate:"required"`
	RetailerFirmState          string `json:"retailer_firm_state" validate:"required"`
	RetailerFirmDistrict       string `json:"retailer_firm_district" validate:"required"`
}

type AdminLoginModel struct {
	AdminID       string `json:"admin_id" validate:"required"`
	AdminPassword string `json:"admin_password" validate:"required,strpwd"`
}

type MasterDistributorLoginModel struct {
	MasterDistributorID       string `json:"master_distributor_id" validate:"required"`
	MasterDistributorPassword string `json:"master_distributor_password" validate:"required,strpwd"`
}

type DistributorLoginModel struct {
	DistributorID       string `json:"distributor_id" validate:"required"`
	DistributorPassword string `json:"distributor_password" validate:"required,strpwd"`
}

type RetailerLoginModel struct {
	RetailerID       string `json:"retailer_id" validate:"required"`
	RetailerPassword string `json:"retailer_password" validate:"required,strpwd"`
}
