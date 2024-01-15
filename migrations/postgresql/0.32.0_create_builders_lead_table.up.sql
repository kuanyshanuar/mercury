CREATE TABLE builder_leads (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    builder_id   BIGINT REFERENCES users (id) UNIQUE,
    status_id    BIGINT REFERENCES lead_statuses(id) NOT NULL DEFAULT 1,
    issued_at    BIGINT NOT NULL,
    expires_at   BIGINT NOT NULL
);