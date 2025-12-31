CREATE TABLE
    IF NOT EXISTS retailers (
        distributor_id TEXT NOT NULL,
        retailer_id TEXT UNIQUE NOT NULL DEFAULT ('MD' || LPAD(nextval('retailer_id_sequence')::TEXT, 9, '0')),
        retailer_name TEXT NOT NULL,
        retailer_father_or_spouse_name TEXT NOT NULL DEFAULT '',
        retailer_email TEXT NOT NULL,
        retailer_phone TEXT NOT NULL,
        retailer_password TEXT NOT NULL,
        retailer_pan_number TEXT NOT NULL,
        retailer_aadhar_number TEXT NOT NULL,
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