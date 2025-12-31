CREATE TABLE
    IF NOT EXISTS admins (
        admin_id TEXT UNIQUE NOT NULL DEFAULT ('A' || LPAD(nextval('admin_id_sequence')::TEXT, 9, '0')),
        admin_name TEXT NOT NULL,
        admin_email TEXT NOT NULL,
        admin_phone TEXT NOT NULL,
        admin_password TEXT NOT NULL,
        admin_wallet NUMERIC(20,2) DEFAULT 0.0,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TRIGGER trg_admins_updated_at BEFORE UPDATE ON admins
			FOR EACH ROW EXECUTE FUNCTION set_updated_at();