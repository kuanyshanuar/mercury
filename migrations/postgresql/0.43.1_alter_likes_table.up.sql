ALTER TABLE likes
    DROP CONSTRAINT likes_article_id_fkey,
    ADD CONSTRAINT likes_article_id_fkey
        FOREIGN KEY (article_id)
            REFERENCES articles(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;