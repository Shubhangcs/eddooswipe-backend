CREATE TABLE
    IF NOT EXISTS master_distributors (
        admin_id TEXT NOT NULL,
        master_distributor_id TEXT UNIQUE NOT NULL DEFAULT ('MD' || LPAD(nextval('master_distributor_id_sequence')::TEXT, 9, '0')),
        master_distributor_name TEXT NOT NULL,
        master_distributor_father_or_spouse_name TEXT NOT NULL DEFAULT '',
        master_distributor_email TEXT NOT NULL,
        master_distributor_phone TEXT NOT NULL,
        master_distributor_password TEXT NOT NULL,
        master_distributor_pan_number TEXT NOT NULL,
        master_distributor_aadhar_number TEXT NOT NULL,
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