CREATE TABLE IF NOT EXISTS residences_bookmarks (
    user_id INT REFERENCES users (id),
    residence_id INT REFERENCES residences (id),
    PRIMARY KEY (user_id, residence_id)
);