package database

import (
	"context"
	"fmt"

	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

func (db *Database) GetAllAdminsQuery(ctx context.Context) (*[]models.GetAdminModel, error) {
	query := `
		SELECT admin_id, admin_name, admin_email,
		admin_phone, admin_wallet::TEXT, created_at::TEXT, updated_at::TEXT
		FROM admins;
	`
	res, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin data")
	}
	defer res.Close()

	var admins []models.GetAdminModel
	for res.Next() {
		var admin models.GetAdminModel
		if err := res.Scan(
			&admin.AdminID,
			&admin.AdminName,
			&admin.AdminEmail,
			&admin.AdminPhone,
			&admin.AdminWalletBalance,
			&admin.CreatedAt,
			&admin.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get admin data")
		}
		admins = append(admins, admin)
	}

	if res.Err() != nil {
		return nil, fmt.Errorf("failed to get admin data")
	}
	return &admins, nil
}

func (db *Database) GetAllMasterDistributorsQuery(ctx context.Context) (*[]models.GetMasterDistributorModel, error) {

	query := `
		SELECT
			master_distributor_id,
			master_distributor_name,
			master_distributor_father_or_spouse_name,
			master_distributor_email,
			master_distributor_phone,
			master_distributor_pan_number,
			master_distributor_aadhar_number,
			master_distributor_gst_number,
			master_distributor_pin_code,
			master_distributor_city,
			master_distributor_state,
			master_distributor_address,
			master_distributor_firm_name,
			master_distributor_firm_address,
			master_distributor_firm_city,
			master_distributor_firm_pin,
			master_distributor_firm_state,
			master_distributor_firm_district,
			master_distributor_added_by,
			master_distributor_added_by_id,
			master_distributor_kyc_status,
			master_distributor_wallet::TEXT,
			is_master_distributor_blocked,
			created_at::TEXT AS created_at,
			updated_at::TEXT AS updated_at
		FROM master_distributors;
	`

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get master distributors")
	}
	defer rows.Close()

	var masterDistributors []models.GetMasterDistributorModel

	for rows.Next() {
		var md models.GetMasterDistributorModel

		if err := rows.Scan(
			&md.MasterDistributorID,
			&md.MasterDistributorName,
			&md.MasterDistributorFatherOrSpouseName,
			&md.MasterDistributorEmail,
			&md.MasterDistributorPhone,
			&md.MasterDistributorPanNumber,
			&md.MasterDistributorAadharNumber,
			&md.MasterDistributorGSTNumber,
			&md.MasterDistributorPinCode,
			&md.MasterDistributorCity,
			&md.MasterDistributorState,
			&md.MasterDistributorAddress,
			&md.MasterDistributorFirmName,
			&md.MasterDistributorFirmAddress,
			&md.MasterDistributorFirmCity,
			&md.MasterDistributorFirmPin,
			&md.MasterDistributorFirmState,
			&md.MasterDistributorFirmDistrict,
			&md.MasterDistributorAddedBy,
			&md.MasterDistributorAddedByID,
			&md.MasterDistributiorKYCStatus,
			&md.MasterDistributorWalletBalance,
			&md.IsMasterDistributorBlocked,
			&md.CreatedAt,
			&md.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get master distributors")
		}

		masterDistributors = append(masterDistributors, md)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get master distributors")
	}

	return &masterDistributors, nil
}

func (db *Database) GetAllDistributorsQuery(ctx context.Context) (*[]models.GetDistributorModel, error) {
	query := `
		SELECT
			distributor_id,
			distributor_name,
			distributor_father_or_spouse_name,
			distributor_email,
			distributor_phone,
			distributor_pan_number,
			distributor_aadhar_number,
			distributor_gst_number,
			distributor_pin_code,
			distributor_city,
			distributor_state,
			distributor_address,
			distributor_firm_name,
			distributor_firm_address,
			distributor_firm_city,
			distributor_firm_pin,
			distributor_firm_state,
			distributor_firm_district,
			distributor_added_by,
			distributor_added_by_id,
			distributor_kyc_status,
			distributor_wallet::TEXT,
			is_distributor_blocked,
			created_at::TEXT AS created_at,
			updated_at::TEXT AS updated_at
		FROM distributors;
	`

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get distributors")
	}
	defer rows.Close()

	var distributors []models.GetDistributorModel

	for rows.Next() {
		var d models.GetDistributorModel

		if err := rows.Scan(
			&d.DistributorID,
			&d.DistributorName,
			&d.DistributorFatherOrSpouseName,
			&d.DistributorEmail,
			&d.DistributorPhone,
			&d.DistributorPanNumber,
			&d.DistributorAadharNumber,
			&d.DistributorGSTNumber,
			&d.DistributorPinCode,
			&d.DistributorCity,
			&d.DistributorState,
			&d.DistributorAddress,
			&d.DistributorFirmName,
			&d.DistributorFirmAddress,
			&d.DistributorFirmCity,
			&d.DistributorFirmPin,
			&d.DistributorFirmState,
			&d.DistributorFirmDistrict,
			&d.DistributorAddedBy,
			&d.DistributorAddedByID,
			&d.DistributorKYCStatus,
			&d.DistributorWalletBalance,
			&d.IsDistributorBlocked,
			&d.CreatedAt,
			&d.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get distributors")
		}

		distributors = append(distributors, d)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get distributors")
	}

	return &distributors, nil
}

func (db *Database) GetAllRetailersQuery(ctx context.Context) (*[]models.GetRetailerModel, error) {
	query := `
		SELECT
			retailer_id,
			retailer_name,
			retailer_father_or_spouse_name,
			retailer_email,
			retailer_phone,
			retailer_pan_number,
			retailer_aadhar_number,
			retailer_gst_number,
			retailer_pin_code,
			retailer_city,
			retailer_state,
			retailer_address,
			retailer_firm_name,
			retailer_firm_address,
			retailer_firm_city,
			retailer_firm_pin,
			retailer_firm_state,
			retailer_firm_district,
			retailer_added_by,
			retailer_added_by_id,
			retailer_kyc_status,
			retailer_wallet::TEXT,
			is_retailer_blocked,
			created_at::TEXT AS created_at,
			updated_at::TEXT AS updated_at
		FROM retailers;
	`

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get retailers")
	}
	defer rows.Close()

	var retailers []models.GetRetailerModel

	for rows.Next() {
		var r models.GetRetailerModel

		if err := rows.Scan(
			&r.RetailerID,
			&r.RetailerName,
			&r.RetailerFatherOrSpouseName,
			&r.RetailerEmail,
			&r.RetailerPhone,
			&r.RetailerPanNumber,
			&r.RetailerAadharNumber,
			&r.RetailerGSTNumber,
			&r.RetailerPinCode,
			&r.RetailerCity,
			&r.RetailerState,
			&r.RetailerAddress,
			&r.RetailerFirmName,
			&r.RetailerFirmAddress,
			&r.RetailerFirmCity,
			&r.RetailerFirmPin,
			&r.RetailerFirmState,
			&r.RetailerFirmDistrict,
			&r.RetailerAddedBy,
			&r.RetailerAddedByID,
			&r.RetailerKYCStatus,
			&r.RetailerWalletBalance,
			&r.IsRetailerBlocked,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to get retailers")
		}

		retailers = append(retailers, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get retailers")
	}

	return &retailers, nil
}
