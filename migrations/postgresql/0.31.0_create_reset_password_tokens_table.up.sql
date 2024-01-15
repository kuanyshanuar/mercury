CREATE TABLE reset_password_tokens (
    user_id INT REFERENCES users (id),
    token VARCHAR,
    confirmed BOOLEAN DEFAULT FALSE,
    created_at BIGINT,
    updated_at BIGINT
);