CREATE TABLE IF NOT EXISTS residences_feedbacks (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    residence_id INT REFERENCES residences (id),
    user_id INT REFERENCES users (id),
    content varchar NOT NULL
);
