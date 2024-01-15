CREATE TABLE IF NOT EXISTS articles (
    id BIGSERIAL NOT NULL PRIMARY KEY ,
    title VARCHAR NOT NULL,
    short_description VARCHAR ,
    slug VARCHAR,
    content VARCHAR,
    views_count BIGINT,
    source_url VARCHAR,
    author_name VARCHAR,
    images VARCHAR[],
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT,
    created_by BIGINT,
    updated_by BIGINT,
    deleted_by BIGINT
);