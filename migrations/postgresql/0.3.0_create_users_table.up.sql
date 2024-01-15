CREATE TABLE IF NOT EXISTS users  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    role_id int REFERENCES roles (id),
    city_id int REFERENCES cities (id),
    first_name VARCHAR,
    last_name VARCHAR,
    email VARCHAR UNIQUE NOT NULL,
    phone VARCHAR,
    password VARCHAR,
    birthdate DATE,
    gender VARCHAR,
    is_verified BOOLEAN DEFAULT FALSE,
    is_banned BOOLEAN DEFAULT FALSE,
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT DEFAULT 0
);

CREATE INDEX idx_users_email ON users (email);