CREATE TABLE IF NOT EXISTS article_views(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    article_id INT REFERENCES articles(id),
    user_id BIGINT
);

CREATE index article_views_index ON article_views(article_id);