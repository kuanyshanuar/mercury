CREATE TABLE IF NOT EXISTS cottage_leads(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    cottage_id INT UNIQUE references cottages(id),
    status_id INT references lead_statuses(id) DEFAULT 1,
    issued_at BIGINT,
    expires_at BIGINT,
    deleted_at INT
);


