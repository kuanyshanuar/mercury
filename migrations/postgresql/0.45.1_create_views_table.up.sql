ALTER TABLE article_views
    DROP CONSTRAINT article_views_article_id_fkey,
    ADD CONSTRAINT article_views_article_id_fkey
        FOREIGN KEY (article_id)
            REFERENCES articles(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;