ALTER TABLE users
    DROP COLUMN IF EXISTS city_id;

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS consultation_phone_number VARCHAR;