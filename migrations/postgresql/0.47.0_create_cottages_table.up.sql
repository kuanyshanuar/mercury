CREATE TABLE IF NOT EXISTS cottages(
	id BIGSERIAL PRIMARY KEY NOT NULL,
    title VARCHAR,
    description VARCHAR,
    address VARCHAR,
    latitude FLOAT8 NOT NULL DEFAULT 0.0,
    longitude FLOAT8 NOT NULL DEFAULT 0.0,
    territory VARCHAR,
	ceiling_height_min FLOAT8 NOT NULL DEFAULT 0.0,
    ceiling_height_max FLOAT8 NOT NULL DEFAULT 0.0,
    building_area FLOAT8 NOT NULL DEFAULT 0.0,
    area_min float8 NOT NULL DEFAULT 0.0,
    area_max float8 NOT NULL DEFAULT 0.0,
    house_amount BIGINT,
    floors_count BIGINT,
    facade VARCHAR,
    rooms_min BIGINT,
    rooms_max BIGINT,
    can_replan bool,
    min_flat_square FLOAT8 NOT NULL DEFAULT 0.0,
	max_flat_square FLOAT8 NOT NULL DEFAULT 0.0,
	price_per_square_min FLOAT8 NOT NULL DEFAULT 0.0,
    price_per_square_max FLOAT8 NOT NULL DEFAULT 0.0,
    city_id INT references cities(id),
    user_id INT references users(id),
    status_id INT references statuses(id),
	district_id INT references districts(id),
    images VARCHAR[],
    housing_class_id INT references housing_classes(id),
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT
);

CREATE INDEX idx_cottages_title ON cottages (title);

CREATE INDEX idx_cottages_district_id ON cottages (district_id);

CREATE INDEX idx_cottages_user ON cottages (user_id);

CREATE INDEX idx_cottages_housing_class ON cottages (housing_class_id);

CREATE INDEX idx_cottages_deleted_at ON cottages (deleted_at);

CREATE TABLE IF NOT EXISTS house_plans(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    cottage_id INT references cottages(id),
    title VARCHAR,
    number_of_rooms INT,
    area FLOAT8 NOT NULL DEFAULT 0.0,
    longitude FLOAT8 NOT NULL DEFAULT 0.0,
    territory FLOAT8 NOT NULL DEFAULT 0.0,
    ceiling_height FLOAT8 NOT NULL DEFAULT 0.0,
    price INT NOT NULL DEFAULT 0,
    price_per_square FLOAT8 NOT NULL DEFAULT 0.0,
    plan_images VARCHAR[],
    house_images VARCHAR[],
    housing_class_id INT references housing_classes(id),
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT
);



