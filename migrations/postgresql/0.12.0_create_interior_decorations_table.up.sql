CREATE TABLE IF NOT EXISTS interior_decorations (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE INDEX idx_interior_decorations_name ON interior_decorations (name);