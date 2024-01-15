CREATE TABLE IF NOT EXISTS dislikes(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    article_id INT REFERENCES articles(id),
    user_id INT references users(id)
);
