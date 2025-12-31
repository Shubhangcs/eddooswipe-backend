CREATE TABLE
    IF NOT EXISTS distributors (
        master_distributor_id TEXT NOT NULL,
        distributor_id TEXT UNIQUE NOT NULL DEFAULT ('D' || LPAD(nextval('distributor_id_sequence')::TEXT, 9, '0')),
        distributor_name TEXT NOT NULL,
        distributor_father_or_spouse_name TEXT NOT NULL DEFAULT '',
        distributor_email TEXT NOT NULL,
        distributor_phone TEXT NOT NULL,
        distributor_password TEXT NOT NULL,
        distributor_pan_number TEXT NOT NULL,
        distributor_aadhar_number TEXT NOT NULL,
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