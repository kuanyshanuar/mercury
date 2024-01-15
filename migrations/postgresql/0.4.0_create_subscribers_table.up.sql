CREATE TABLE IF NOT EXISTS subscribers (
    id BIGINT NOT NULL,
    builder_id INT REFERENCES users (id),
    subscriber_id INT REFERENCES users (id),
    PRIMARY KEY (builder_id, subscriber_id)
);
