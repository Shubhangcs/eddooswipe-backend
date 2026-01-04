CREATE TABLE
    IF NOT EXISTS ledger_entries (
        ledger_transaction_id TEXT PRIMARY KEY NOT NULL DEFAULT ('LT' || LPAD(nextval('ledger_transaction_id_sequence')::TEXT, 9, '0')),
        transactor_id TEXT NOT NULL,
        reference_id TEXT NOT NULL,
        remarks TEXT NOT NULL,
        credit_amount TEXT NOT NULL DEFAULT '',
        debit_amount TEXT NOT NULL DEFAULT '',
        latest_balance TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TABLE 
    IF NOT EXISTS operators (
        operator_id TEXT NOT NULL,
        operator_name TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT NOW()
    );

CREATE TABLE 
    IF NOT EXISTS commisions (
        commision_id TEXT PRIMARY KEY NOT NULL,
        operator_id TEXT NOT NULL,
        operator_name TEXT NOT NULL,
        slab_start NUMERIC(20,2) DEFAULT 0,
        slab_end NUMERIC(20,2) DEFAULT 0,
        total_commision NUMERIC(20,0) DEFAULT 0,
        admin_commision NUMERIC(20,2) DEFAULT 0,
        master_distributor_commision NUMERIC(20,2) DEFAULT 0,
        distributor_commision NUMERIC(20,2) DEFAULT 0,
        retailer_commision NUMERIC(20,2) DEFAULT 0,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        FOREIGN KEY (operator_id) REFERENCES operators(operator_id) ON DELETE CASCADE
    );

