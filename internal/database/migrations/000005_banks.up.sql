CREATE TABLE
    IF NOT EXISTS admin_banks (
        admin_id TEXT NOT NULL,
        admin_bank_name TEXT NOT NULL,
        admin_bank_address TEXT NOT NULL,
        admin_bank_account_holder_name TEXT NOT NULL,
        admin_bank_ifsc_code TEXT NOT NULL,
        admin_bank_account_number TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        FOREIGN KEY (admin_id) REFERENCES admins(admin_id) ON DELETE CASCADE
    );

CREATE TRIGGER trg_admin_banks_updated_at BEFORE UPDATE ON admin_banks
		FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE
    IF NOT EXISTS retailer_banks (
        retailer_id TEXT NOT NULL,
        retailer_bank_name TEXT NOT NULL,
        retailer_bank_address TEXT NOT NULL,
        retailer_bank_account_holder_name TEXT NOT NULL,
        retailer_bank_ifsc_code TEXT NOT NULL,
        retailer_bank_account_number TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        FOREIGN KEY (retailer_id) REFERENCES retailers(retailer_id) ON DELETE CASCADE
    );

CREATE TRIGGER trg_retailer_banks_updated_at BEFORE UPDATE ON retailer_banks
		FOR EACH ROW EXECUTE FUNCTION set_updated_at();