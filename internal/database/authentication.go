package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

func (db *Database) CreateAdminQuery(ctx context.Context, req models.CreateAdminModel) error {
	query := `
		INSERT INTO admins(
			admin_name,
			admin_email,
			admin_phone,
			admin_password
		) VALUES(
			@admin_name,
			@admin_email,
			@admin_phone,
			@admin_password 
		);
	`
	hash, err := pkg.GenerateHashedPassword(req.AdminPassword)
	if err != nil {
		return err
	}
	_, err = db.pool.Exec(ctx, query, pgx.NamedArgs{
		"admin_name":     req.AdminName,
		"admin_email":    req.AdminEmail,
		"admin_phone":    req.AdminPhone,
		"admin_password": hash,
	})
	if err != nil {
		return fmt.Errorf("failed to create admin")
	}
	return nil
}

func (db *Database) CreateMasterDistributorQuery(ctx context.Context, req models.CreateMasterDistributorModel) error {
	query := `
		WITH admin_details AS (
			SELECT admin_name FROM admins WHERE admin_id=@admin_id
		)
		INSERT INTO master_distributors(
			admin_id,
			master_distributor_name,
			master_distributor_father_or_spouse_name,
			master_distributor_email,
			master_distributor_phone,
			master_distributor_password,
			master_distributor_pan_number,
			master_distributor_aadhar_number,
			master_distributor_gst_number,
			master_distributor_pin_code,
			master_distributor_address,
			master_distributor_city,
			master_distributor_state,
			master_distributor_firm_name,
			master_distributor_firm_address,
			master_distributor_firm_city,
			master_distributor_firm_pin,
			master_distributor_firm_state,
			master_distributor_firm_district,
			master_distributor_added_by,
			master_distributor_added_by_id
		) VALUES (
			@admin_id,
			@master_distributor_name,
			@master_distributor_father_or_spouse_name,
			@master_distributor_email,
			@master_distributor_phone,
			@master_distributor_password,
			@master_distributor_pan_number,
			@master_distributor_aadhar_number,
			@master_distributor_gst_number,
			@master_distributor_pin_code,
			@master_distributor_address,
			@master_distributor_city,
			@master_distributor_state,
			@master_distributor_firm_name,
			@master_distributor_firm_address,
			@master_distributor_firm_city,
			@master_distributor_firm_pin,
			@master_distributor_firm_state,
			@master_distributor_firm_district,
			(SELECT admin_name FROM admin_details),
			@admin_id
		);
	`
	hash, err := pkg.GenerateHashedPassword(req.MasterDistributorPassword)
	if err != nil {
		return err
	}
	_, err = db.pool.Exec(ctx, query, pgx.NamedArgs{
		"admin_id":                req.AdminID,
		"master_distributor_name": req.MasterDistributorName,
		"master_distributor_father_or_spouse_name": req.MasterDistributorFatherOrSpouseName,
		"master_distributor_email":                 req.MasterDistributorEmail,
		"master_distributor_phone":                 req.MasterDistributorPhone,
		"master_distributor_password":              hash,
		"master_distributor_pan_number":            req.MasterDistributorPanNumber,
		"master_distributor_aadhar_number":         req.MasterDistributorAadharNumber,
		"master_distributor_gst_number":            req.MasterDistributorGSTNumber,
		"master_distributor_pin_code":              req.MasterDistributorPinCode,
		"master_distributor_address":               req.MasterDistributorAddress,
		"master_distributor_city":                  req.MasterDistributorCity,
		"master_distributor_state":                 req.MasterDistributorState,
		"master_distributor_firm_name":             req.MasterDistributorFirmName,
		"master_distributor_firm_address":          req.MasterDistributorFirmAddress,
		"master_distributor_firm_city":             req.MasterDistributorFirmCity,
		"master_distributor_firm_pin":              req.MasterDistributorFirmPin,
		"master_distributor_firm_state":            req.MasterDistributorFirmState,
		"master_distributor_firm_district":         req.MasterDistributorFirmDistrict,
	})

	if err != nil {
		return fmt.Errorf("failed to create master distributor")
	}
	return nil
}

func (db *Database) CreateDistributorQuery(ctx context.Context, req models.CreateDistributorModel) error {
	query := `
		INSERT INTO distributors (
    		master_distributor_id,
    		distributor_name,
    		distributor_father_or_spouse_name,
    		distributor_email,
    		distributor_phone,
    		distributor_password,
    		distributor_pan_number,
    		distributor_aadhar_number,
    		distributor_gst_number,
    		distributor_pin_code,
    		distributor_address,
    		distributor_city,
    		distributor_state,
    		distributor_firm_name,
    		distributor_firm_address,
    		distributor_firm_city,
    		distributor_firm_pin,
    		distributor_firm_state,
    		distributor_firm_district,
    		distributor_added_by,
    		distributor_added_by_id
		) VALUES (
			@master_distributor_id,
    		@distributor_name,
    		@distributor_father_or_spouse_name,
    		@distributor_email,
    		@distributor_phone,
    		@distributor_password,
    		@distributor_pan_number,
    		@distributor_aadhar_number,
    		@distributor_gst_number,
    		@distributor_pin_code,
    		@distributor_address,
    		@distributor_city,
    		@distributor_state,
    		@distributor_firm_name,
    		@distributor_firm_address,
    		@distributor_firm_city,
    		@distributor_firm_pin,
    		@distributor_firm_state,
    		@distributor_firm_district,
			@distributor_added_by,
			@distributor_added_by_id
		);
	`
	hash, err := pkg.GenerateHashedPassword(req.DistributorPassword)
	if err != nil {
		return err
	}
	_, err = db.pool.Exec(ctx, query, pgx.NamedArgs{
		"master_distributor_id":             req.MasterDistributorID,
		"distributor_name":                  req.DistributorName,
		"distributor_father_or_spouse_name": req.DistributorFatherOrSpouseName,
		"distributor_email":                 req.DistributorEmail,
		"distributor_phone":                 req.DistributorPhone,
		"distributor_password":              hash,
		"distributor_pan_number":            req.DistributorPanNumber,
		"distributor_aadhar_number":         req.DistributorAadharNumber,
		"distributor_gst_number":            req.DistributorGSTNumber,
		"distributor_pin_code":              req.DistributorPinCode,
		"distributor_address":               req.DistributorAddress,
		"distributor_city":                  req.DistributorCity,
		"distributor_state":                 req.DistributorState,
		"distributor_firm_name":             req.DistributorFirmName,
		"distributor_firm_address":          req.DistributorFirmAddress,
		"distributor_firm_city":             req.DistributorFirmCity,
		"distributor_firm_pin":              req.DistributorFirmPin,
		"distributor_firm_state":            req.DistributorFirmState,
		"distributor_firm_district":         req.DistributorFirmDistrict,
		"distributor_added_by":              req.CreatorName,
		"distributor_added_by_id":           req.CreatorID,
	})

	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create distributor")
	}
	return nil
}

func (db *Database) CreateRetailerQuery(ctx context.Context, req models.CreateRetailerModel) error {
	query := `
		INSERT INTO retailers (
    		distributor_id,
    		retailer_name,
    		retailer_father_or_spouse_name,
    		retailer_email,
    		retailer_phone,
    		retailer_password,
    		retailer_pan_number,
    		retailer_aadhar_number,
    		retailer_gst_number,
    		retailer_pin_code,
    		retailer_address,
    		retailer_city,
    		retailer_state,
    		retailer_firm_name,
    		retailer_firm_address,
    		retailer_firm_city,
    		retailer_firm_pin,
    		retailer_firm_state,
    		retailer_firm_district,
    		retailer_added_by,
    		retailer_added_by_id
		) VALUES (
			@distributor_id,
    		@retailer_name,
    		@retailer_father_or_spouse_name,
    		@retailer_email,
    		@retailer_phone,
    		@retailer_password,
    		@retailer_pan_number,
    		@retailer_aadhar_number,
    		@retailer_gst_number,
    		@retailer_pin_code,
    		@retailer_address,
    		@retailer_city,
    		@retailer_state,
    		@retailer_firm_name,
    		@retailer_firm_address,
    		@retailer_firm_city,
    		@retailer_firm_pin,
    		@retailer_firm_state,
    		@retailer_firm_district,
    		@retailer_added_by,
    		@retailer_added_by_id
		);
	`
	hash, err := pkg.GenerateHashedPassword(req.RetailerPassword)
	if err != nil {
		return err
	}
	_, err = db.pool.Exec(ctx, query, pgx.NamedArgs{
		"distributor_id":                 req.DistributorID,
		"retailer_name":                  req.RetailerName,
		"retailer_father_or_spouse_name": req.RetailerFatherOrSpouseName,
		"retailer_email":                 req.RetailerEmail,
		"retailer_phone":                 req.RetailerPhone,
		"retailer_password":              hash,
		"retailer_pan_number":            req.RetailerPanNumber,
		"retailer_aadhar_number":         req.RetailerAadharNumber,
		"retailer_gst_number":            req.RetailerGSTNumber,
		"retailer_pin_code":              req.RetailerPinCode,
		"retailer_address":               req.RetailerAddress,
		"retailer_city":                  req.RetailerCity,
		"retailer_state":                 req.RetailerState,
		"retailer_firm_name":             req.RetailerFirmName,
		"retailer_firm_address":          req.RetailerFirmAddress,
		"retailer_firm_city":             req.RetailerFirmCity,
		"retailer_firm_pin":              req.RetailerFirmPin,
		"retailer_firm_state":            req.RetailerFirmState,
		"retailer_firm_district":         req.RetailerFirmDistrict,
		"retailer_added_by":              req.CreatorName,
		"retailer_added_by_id":           req.CreatorID,
	})

	if err != nil {
		return fmt.Errorf("failed to create retailer")
	}
	return nil
}

func (db *Database) LoginAdminQuery(ctx context.Context, req models.AdminLoginModel) (*models.JWTTokenModel, error) {
	var res struct {
		AdminID       string
		AdminName     string
		AdminPassword string
	}
	query := `
		SELECT admin_id, admin_name, admin_password
		FROM admins
		WHERE admin_id=@admin_id;
	`
	err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"admin_id": req.AdminID,
	}).Scan(
		&res.AdminID,
		&res.AdminName,
		&res.AdminPassword,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user")
	}

	if err := pkg.ComparePassword(res.AdminPassword, req.AdminPassword); err != nil {
		return nil, fmt.Errorf("failed to authenticate incorrect password")
	}

	return &models.JWTTokenModel{
		ID:   res.AdminID,
		Name: res.AdminName,
	}, nil
}

func (db *Database) LoginMasterDistributorQuery(ctx context.Context, req models.MasterDistributorLoginModel) (*models.JWTTokenModel, error) {
	var res struct {
		AdminID                   string
		MasterDistributorID       string
		MasterDistributorName     string
		MasterDistributorPassword string
		IsMDBlocked               bool
	}
	query := `
		SELECT admin_id ,master_distributor_id, master_distributor_name, master_distributor_password, is_master_distributor_blocked
		FROM master_distributors
		WHERE master_distributor_id=@master_distributor_id;
	`
	err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"master_distributor_id": req.MasterDistributorID,
	}).Scan(
		&res.AdminID,
		&res.MasterDistributorID,
		&res.MasterDistributorName,
		&res.MasterDistributorPassword,
		&res.IsMDBlocked,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user")
	}

	if res.IsMDBlocked {
		return nil, fmt.Errorf("failed to login master distributor is blocked")
	}

	if err := pkg.ComparePassword(res.MasterDistributorPassword, req.MasterDistributorPassword); err != nil {
		return nil, fmt.Errorf("failed to authenticate incorrect password")
	}

	return &models.JWTTokenModel{
		AdminID: res.AdminID,
		ID:      res.MasterDistributorID,
		Name:    res.MasterDistributorName,
	}, nil
}

func (db *Database) LoginDistributorQuery(ctx context.Context, req models.DistributorLoginModel) (*models.JWTTokenModel, error) {
	var res struct {
		AdminID              string
		DistributorID        string
		DistributorName      string
		DistributorPassword  string
		IsDistributorBlocked bool
	}
	query := `
		SELECT
    		d.distributor_id,
    		d.distributor_name,
    		d.distributor_password,
    		d.is_distributor_blocked,
    		md.admin_id
		FROM distributors d
		JOIN master_distributors md
    		ON md.master_distributor_id = d.master_distributor_id
		WHERE d.distributor_id = @distributor_id;
	`
	err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"distributor_id": req.DistributorID,
	}).Scan(
		&res.DistributorID,
		&res.DistributorName,
		&res.DistributorPassword,
		&res.IsDistributorBlocked,
		&res.AdminID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user")
	}

	if res.IsDistributorBlocked {
		return nil, fmt.Errorf("failed to login distributor is blocked")
	}

	if err := pkg.ComparePassword(res.DistributorPassword, req.DistributorPassword); err != nil {
		return nil, fmt.Errorf("failed to authenticate incorrect password")
	}

	return &models.JWTTokenModel{
		AdminID: res.AdminID,
		ID:      res.DistributorID,
		Name:    res.DistributorName,
	}, nil
}

func (db *Database) LoginRetailerQuery(ctx context.Context, req models.RetailerLoginModel) (*models.JWTTokenModel, error) {
	var res struct {
		AdminID           string
		RetailerID        string
		RetailerName      string
		RetailerPassword  string
		IsRetailerBlocked bool
	}
	query := `
		SELECT
    		r.retailer_id,
    		r.retailer_name,
    		r.retailer_password,
    		r.is_retailer_blocked,
    		md.admin_id
		FROM retailers r
		JOIN distributors d
    		ON d.distributor_id = r.distributor_id
		JOIN master_distributors md
    		ON md.master_distributor_id = d.master_distributor_id
		WHERE r.retailer_id = @retailer_id;

	`
	err := db.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"retailer_id": req.RetailerID,
	}).Scan(
		&res.RetailerID,
		&res.RetailerName,
		&res.RetailerPassword,
		&res.IsRetailerBlocked,
		&res.AdminID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user")
	}
	if res.IsRetailerBlocked {
		return nil, fmt.Errorf("failed to login retailer is blocked")
	}

	if err := pkg.ComparePassword(res.RetailerPassword, req.RetailerPassword); err != nil {
		return nil, fmt.Errorf("failed to authenticate incorrect password")
	}

	return &models.JWTTokenModel{
		AdminID: res.AdminID,
		ID:      res.RetailerID,
		Name:    res.RetailerName,
	}, nil
}
