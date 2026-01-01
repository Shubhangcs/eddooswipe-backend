CREATE TABLE
    IF NOT EXISTS ledger_entries (
        ledger_transaction_id TEXT UNIQUE NOT NULL DEFAULT ('LT' || LPAD(nextval('ledger_transaction_id_sequence')::TEXT, 9, '0')),
        transactor_id TEXT NOT NULL,
        reference_id TEXT NOT NULL,
        remarks TEXT NOT NULL,
        credit_amount TEXT NOT NULL DEFAULT '',
        debit_amount TEXT NOT NULL DEFAULT '',
        latest_balance TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );