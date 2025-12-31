package models

type GetAdminModel struct {
	AdminID    string `json:"admin_id"`
	AdminName  string `json:"admin_name"`
	AdminEmail string `json:"admin_email"`
	AdminPhone string `json:"admin_phone"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type GetMasterDistributorModel struct {
	MasterDistributorID                 string `json:"master_distributor_id"`
	MasterDistributorName               string `json:"master_distributor_name"`
	MasterDistributorFatherOrSpouseName string `json:"master_distributor_father_or_spouse_name,omitempty"`
	MasterDistributorEmail              string `json:"master_distributor_email"`
	MasterDistributorPhone              string `json:"master_distributor_phone"`
	MasterDistributorPanNumber          string `json:"master_distributor_pan_number"`
	MasterDistributorAadharNumber       string `json:"master_distributor_aadhar_number"`
	MasterDistributorGSTNumber          string `json:"master_distributor_gst_number"`
	MasterDistributorPinCode            string `json:"master_distributor_pin_code"`
	MasterDistributorCity               string `json:"master_distributor_city"`
	MasterDistributorState              string `json:"master_distributor_state"`
	MasterDistributorAddress            string `json:"master_distributor_address"`
	MasterDistributorFirmName           string `json:"master_distributor_firm_name"`
	MasterDistributorFirmAddress        string `json:"master_distributor_firm_address"`
	MasterDistributorFirmCity           string `json:"master_distributor_firm_city"`
	MasterDistributorFirmPin            string `json:"master_distributor_firm_pin"`
	MasterDistributorFirmState          string `json:"master_distributor_firm_state"`
	MasterDistributorFirmDistrict       string `json:"master_distributor_firm_district"`
	MasterDistributorAddedBy            string `json:"master_distributor_added_by"`
	MasterDistributorAddedByID          string `json:"master_distributor_added_by_id"`
	MasterDistributiorKYCStatus         bool   `json:"master_distributor_kyc_status"`
	IsMasterDistributorBlocked          bool   `json:"is_master_distributor_blocked"`
	CreatedAt                           string `json:"created_at"`
	UpdatedAt                           string `json:"updated_at"`
}

type GetDistributorModel struct {
	DistributorID                 string `json:"distributor_id"`
	DistributorName               string `json:"distributor_name"`
	DistributorFatherOrSpouseName string `json:"distributor_father_or_spouse_name,omitempty"`
	DistributorEmail              string `json:"distributor_email"`
	DistributorPhone              string `json:"distributor_phone"`
	DistributorPanNumber          string `json:"distributor_pan_number"`
	DistributorAadharNumber       string `json:"distributor_aadhar_number"`
	DistributorGSTNumber          string `json:"distributor_gst_number"`
	DistributorPinCode            string `json:"distributor_pin_code"`
	DistributorCity               string `json:"distributor_city"`
	DistributorState              string `json:"distributor_state"`
	DistributorAddress            string `json:"distributor_address"`
	DistributorFirmName           string `json:"distributor_firm_name"`
	DistributorFirmAddress        string `json:"distributor_firm_address"`
	DistributorFirmCity           string `json:"distributor_firm_city"`
	DistributorFirmPin            string `json:"distributor_firm_pin"`
	DistributorFirmState          string `json:"distributor_firm_state"`
	DistributorFirmDistrict       string `json:"distributor_firm_district"`
	DistributorAddedBy            string `json:"distributor_added_by"`
	DistributorAddedByID          string `json:"distributor_added_by_id"`
	DistributorKYCStatus         bool   `json:"distributor_kyc_status"`
	IsDistributorBlocked          bool   `json:"is_distributor_blocked"`
	CreatedAt                     string `json:"created_at"`
	UpdatedAt                     string `json:"updated_at"`
}

type GetRetailerModel struct {
	RetailerID                 string `json:"retailer_id"`
	RetailerName               string `json:"retailer_name"`
	RetailerFatherOrSpouseName string `json:"retailer_father_or_spouse_name,omitempty"`
	RetailerEmail              string `json:"retailer_email"`
	RetailerPhone              string `json:"retailer_phone"`
	RetailerPanNumber          string `json:"retailer_pan_number"`
	RetailerAadharNumber       string `json:"retailer_aadhar_number"`
	RetailerGSTNumber          string `json:"retailer_gst_number"`
	RetailerPinCode            string `json:"retailer_pin_code"`
	RetailerCity               string `json:"retailer_city"`
	RetailerState              string `json:"retailer_state"`
	RetailerAddress            string `json:"retailer_address"`
	RetailerFirmName           string `json:"retailer_firm_name"`
	RetailerFirmAddress        string `json:"retailer_firm_address"`
	RetailerFirmCity           string `json:"retailer_firm_city"`
	RetailerFirmPin            string `json:"retailer_firm_pin"`
	RetailerFirmState          string `json:"retailer_firm_state"`
	RetailerFirmDistrict       string `json:"retailer_firm_district"`
	RetailerAddedBy            string `json:"retailer_added_by"`
	RetailerAddedByID          string `json:"retailer_added_by_id"`
	RetailerKYCStatus          bool   `json:"retailer_kyc_status"`
	IsRetailerBlocked          bool   `json:"is_retailer_blocked"`
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
}
