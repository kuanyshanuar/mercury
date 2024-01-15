CREATE TABLE IF NOT EXISTS residences  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT REFERENCES users (id),
    status_id INT REFERENCES statuses (id),
    city_id INT REFERENCES cities (id),
    housing_class_id INT REFERENCES housing_classes (id),
    construction_type_id INT REFERENCES construction_types (id),
    title VARCHAR NOT NULL,
    description VARCHAR,
    district_id INT REFERENCES districts (id),
    address VARCHAR,
    latitude FLOAT8,
    longitude FLOAT8,
    deadline_year INT,
    deadline_quarter INT NOT NULL DEFAULT 0,
    floors_min INT NOT NULL DEFAULT 0,
    floors_max INT NOT NULL DEFAULT 0,
    rooms_min INT NOT NULL DEFAULT 0,
    rooms_max INT NOT NULL DEFAULT 0,
    ceiling_height FLOAT8,
    flats_count INT,
    has_hgf BOOLEAN,
    has_elevator BOOLEAN NOT NULL DEFAULT FALSE,
    price_per_square_min INT NOT NULL DEFAULT 0,
    price_min INT NOT NULL DEFAULT 0,
    price_max INT NOT NULL DEFAULT 0,
    area_min FLOAT8 NOT NULL DEFAULT 0.0,
    area_max FLOAT8 NOT NULL DEFAULT 0.0,
    comments_count FLOAT8,
    title_image VARCHAR,
    images VARCHAR[],
    views INT NOT NULL DEFAULT 0,
    likes INT NOT NULL DEFAULT 0,
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT
);

CREATE INDEX idx_residences_title ON residences (title);

CREATE INDEX idx_residences_district_id ON residences (district_id);

CREATE INDEX idx_residences_user ON residences (user_id);

CREATE INDEX idx_residences_housing_class ON residences (housing_class_id);

CREATE INDEX idx_residences_deleted_at ON residences (deleted_at);
