CREATE TABLE IF NOT EXISTS cottages_bookmarks (
                                                    user_id INT REFERENCES users (id),
                                                    cottage_id INT REFERENCES cottages (id),
                                                    created_at BIGINT,
                                                    PRIMARY KEY (user_id, cottage_id)
);