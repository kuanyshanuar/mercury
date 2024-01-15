CREATE TABLE residences_contact_details (
    id           BIGSERIAL NOT NULL PRIMARY KEY,
    residence_id BIGINT    NOT NULL REFERENCES residences(id),
    full_name    VARCHAR   NOT NULL,
    phone        VARCHAR   NOT NULL,
    has_accepted BOOL      NOT NULL DEFAULT TRUE
);