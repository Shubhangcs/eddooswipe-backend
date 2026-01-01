DROP TABLE IF EXISTS admins;

DROP TRIGGER IF EXISTS trg_admins_updated_at ON admins;

DROP TABLE IF EXISTS master_distributors;

DROP TRIGGER IF EXISTS trg_master_distributors_updated_at ON master_distributors;

DROP TABLE IF EXISTS distributors;

DROP TRIGGER IF EXISTS trg_distributors_updated_at ON distributors;

DROP TABLE IF EXISTS retailers;

DROP TRIGGER IF EXISTS trg_retailers_updated_at ON retailers;