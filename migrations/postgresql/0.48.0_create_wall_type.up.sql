CREATE TABLE IF NOT EXISTS wall_types(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS warming_types(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS window_types(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    name VARCHAR
);