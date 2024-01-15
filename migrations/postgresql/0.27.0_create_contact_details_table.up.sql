CREATE TABLE contact_details (
    id           BIGSERIAL NOT NULL PRIMARY KEY,
    full_name    VARCHAR   NOT NULL,
    email        VARCHAR   NOT NULL,
    phone        VARCHAR   NOT NULL,
    message      VARCHAR   NOT NULL,
    has_accepted BOOL      NOT NULL DEFAULT TRUE
);