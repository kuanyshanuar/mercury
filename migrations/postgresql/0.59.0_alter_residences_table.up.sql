ALTER TABLE residences ADD COLUMN IF NOT EXISTS sale_status_id BIGINT REFERENCES sale_statuses(id);