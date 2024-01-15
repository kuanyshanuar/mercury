CREATE TYPE crud_type AS ENUM ('C', 'R', 'U', 'D');

CREATE TABLE IF NOT EXISTS permissions (
    id            BIGSERIAL NOT NULL PRIMARY KEY,
    endpoint_name VARCHAR   NOT NULL UNIQUE,
    action        VARCHAR   NOT NULL,
    crud_type     crud_type NOT NULL,
    is_active     BOOLEAN   NOT NULL DEFAULT FALSE,
    created_at    BIGINT,
    updated_at    BIGINT
);
