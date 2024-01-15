CREATE TABLE IF NOT EXISTS sms_messages (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    phone VARCHAR NOT NULL,
    message VARCHAR NOT NULL,
    expires_at BIGINT,
    created_at BIGINT
);

CREATE INDEX idx_sms_messages_phone ON sms_messages (phone);