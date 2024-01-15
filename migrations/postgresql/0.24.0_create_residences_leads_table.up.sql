CREATE TABLE residence_leads (
    residence_id BIGINT REFERENCES residences (id),
    issued_at    BIGINT,
    expires_at   BIGINT
);