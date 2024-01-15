ALTER TABLE dislikes
    DROP CONSTRAINT dislikes_article_id_fkey,
    ADD CONSTRAINT dislikes_article_id_fkey
        FOREIGN KEY (article_id)
            REFERENCES articles(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;