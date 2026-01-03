CREATE TABLE
    IF NOT EXISTS fund_requests (
        fund_request_id TEXT UNIQUE NOT NULL DEFAULT ('FR' || LPAD(nextval('fund_request_id_sequence')::TEXT, 9, '0')),
        requester_id TEXT NOT NULL,
        requester_name TEXT NOT NULL,
        request_to_id TEXT NOT NULL,
        request_to_name TEXT NOT NULL,
        payment_mode TEXT NOT NULL,
        deposit_date TEXT NOT NULL,
        amount NUMERIC(20, 2) NOT NULL,
        account_number TEXT DEFAULT '',
        bank_name TEXT DEFAULT '',
        utr_number TEXT DEFAULT '',
        collection_person TEXT DEFAULT '',
        fund_request_status TEXT NOT NULL CHECK (
            fund_request_status IN ('PENDING', 'ACCEPTED', 'REJECTED')
        ),
        requester_remarks TEXT NOT NULL DEFAULT '',
        request_to_remarks TEXT NOT NULL DEFAULT '',
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TRIGGER trg_fund_request_updated_at BEFORE UPDATE ON fund_requests
		FOR EACH ROW EXECUTE FUNCTION set_updated_at();