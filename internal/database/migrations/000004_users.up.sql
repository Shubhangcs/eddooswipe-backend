CREATE TABLE
    IF NOT EXISTS admins (
        admin_id TEXT PRIMARY KEY NOT NULL DEFAULT ('A' || LPAD(nextval('admin_id_sequence')::TEXT, 9, '0')),
        admin_name TEXT NOT NULL,
        admin_email TEXT NOT NULL UNIQUE,
        admin_phone TEXT NOT NULL UNIQUE,
        admin_password TEXT NOT NULL,
        admin_wallet NUMERIC(20,2) DEFAULT 0.0,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TRIGGER trg_admins_updated_at BEFORE UPDATE ON admins
			FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE
    IF NOT EXISTS master_distributors (
        admin_id TEXT NOT NULL,
        master_distributor_id TEXT PRIMARY KEY NOT NULL DEFAULT ('MD' || LPAD(nextval('master_distributor_id_sequence')::TEXT, 9, '0')),
        master_distributor_name TEXT NOT NULL,
        master_distributor_father_or_spouse_name TEXT NOT NULL DEFAULT '',
        master_distributor_email TEXT NOT NULL UNIQUE,
        master_distributor_phone TEXT NOT NULL UNIQUE,
        master_distributor_password TEXT NOT NULL,
        master_distributor_pan_number TEXT NOT NULL UNIQUE,
        master_distributor_aadhar_number TEXT NOT NULL UNIQUE,
        master_distributor_gst_number TEXT NOT NULL DEFAULT '',
        master_distributor_pin_code TEXT NOT NULL,
        master_distributor_address TEXT NOT NULL,
        master_distributor_city TEXT NOT NULL,
        master_distributor_state TEXT NOT NULL,
        master_distributor_firm_name TEXT NOT NULL,
        master_distributor_firm_address TEXT NOT NULL,
        master_distributor_firm_city TEXT NOT NULL,
        master_distributor_firm_pin TEXT NOT NULL,
        master_distributor_firm_state TEXT NOT NULL,
        master_distributor_firm_district TEXT NOT NULL,
        master_distributor_wallet NUMERIC(20,2) DEFAULT 0.0,
        master_distributor_kyc_status BOOLEAN DEFAULT FALSE,
        master_distributor_added_by TEXT NOT NULL,
        master_distributor_added_by_id TEXT NOT NULL,
        is_master_distributor_blocked BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        FOREIGN KEY (admin_id) REFERENCES admins(admin_id) ON DELETE CASCADE
    );

CREATE TRIGGER trg_master_distributors_updated_at BEFORE UPDATE ON master_distributors
			FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE
    IF NOT EXISTS distributors (
        master_distributor_id TEXT NOT NULL,
        distributor_id TEXT PRIMARY KEY NOT NULL DEFAULT ('D' || LPAD(nextval('distributor_id_sequence')::TEXT, 9, '0')),
        distributor_name TEXT NOT NULL,
        distributor_father_or_spouse_name TEXT NOT NULL DEFAULT '',
        distributor_email TEXT NOT NULL UNIQUE,
        distributor_phone TEXT NOT NULL UNIQUE,
        distributor_password TEXT NOT NULL,
        distributor_pan_number TEXT NOT NULL UNIQUE,
        distributor_aadhar_number TEXT NOT NULL UNIQUE,
        distributor_gst_number TEXT NOT NULL DEFAULT '',
        distributor_pin_code TEXT NOT NULL,
        distributor_address TEXT NOT NULL,
        distributor_city TEXT NOT NULL,
        distributor_state TEXT NOT NULL,
        distributor_firm_name TEXT NOT NULL,
        distributor_firm_address TEXT NOT NULL,
        distributor_firm_city TEXT NOT NULL,
        distributor_firm_pin TEXT NOT NULL,
        distributor_firm_state TEXT NOT NULL,
        distributor_firm_district TEXT NOT NULL,
        distributor_wallet NUMERIC(20,2) DEFAULT 0.0,
        distributor_kyc_status BOOLEAN DEFAULT FALSE,
        distributor_added_by TEXT NOT NULL,
        distributor_added_by_id TEXT NOT NULL,
        is_distributor_blocked BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        FOREIGN KEY (master_distributor_id) REFERENCES master_distributors(master_distributor_id) ON DELETE CASCADE
    );

CREATE TRIGGER trg_distributors_updated_at BEFORE UPDATE ON distributors
			FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE
    IF NOT EXISTS retailers (
        distributor_id TEXT NOT NULL,
        retailer_id TEXT PRIMARY KEY NOT NULL DEFAULT ('R' || LPAD(nextval('retailer_id_sequence')::TEXT, 9, '0')),
        retailer_name TEXT NOT NULL,
        retailer_father_or_spouse_name TEXT NOT NULL DEFAULT '',
        retailer_email TEXT NOT NULL UNIQUE,
        retailer_phone TEXT NOT NULL UNIQUE,
        retailer_password TEXT NOT NULL,
        retailer_pan_number TEXT NOT NULL UNIQUE,
        retailer_aadhar_number TEXT NOT NULL UNIQUE,
        retailer_gst_number TEXT NOT NULL DEFAULT '',
        retailer_pin_code TEXT NOT NULL,
        retailer_address TEXT NOT NULL,
        retailer_city TEXT NOT NULL,
        retailer_state TEXT NOT NULL,
        retailer_firm_name TEXT NOT NULL,
        retailer_firm_address TEXT NOT NULL,
        retailer_firm_city TEXT NOT NULL,
        retailer_firm_pin TEXT NOT NULL,
        retailer_firm_state TEXT NOT NULL,
        retailer_firm_district TEXT NOT NULL,
        retailer_wallet NUMERIC(20,2) DEFAULT 0.0,
        retailer_kyc_status BOOLEAN DEFAULT FALSE,
        retailer_added_by TEXT NOT NULL,
        retailer_added_by_id TEXT NOT NULL,
        is_retailer_blocked BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        FOREIGN KEY (distributor_id) REFERENCES distributors(distributor_id) ON DELETE CASCADE
    );

CREATE TRIGGER trg_retailers_updated_at BEFORE UPDATE ON retailers
			FOR EACH ROW EXECUTE FUNCTION set_updated_at();