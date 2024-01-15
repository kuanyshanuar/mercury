ALTER TABLE residences
    DROP CONSTRAINT residences_city_id_fkey,
    ADD CONSTRAINT residences_city_id_fkey
        FOREIGN KEY (city_id)
            REFERENCES cities(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE residences
    DROP CONSTRAINT residences_construction_type_id_fkey,
    ADD CONSTRAINT residences_construction_type_id_fkey
        FOREIGN KEY (construction_type_id)
            REFERENCES construction_types(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE residences
    DROP CONSTRAINT residences_district_id_fkey,
    ADD CONSTRAINT residences_district_id_fkey
        FOREIGN KEY (district_id)
            REFERENCES districts(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE residences
    DROP CONSTRAINT residences_housing_class_id_fkey,
    ADD CONSTRAINT residences_housing_class_id_fkey
        FOREIGN KEY (housing_class_id)
            REFERENCES housing_classes(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE residences
    DROP CONSTRAINT residences_status_id_fkey,
    ADD CONSTRAINT residences_status_id_fkey
        FOREIGN KEY (status_id)
        REFERENCES statuses(id)
        ON UPDATE CASCADE
        ON DELETE SET NULL;

ALTER TABLE residences
    DROP CONSTRAINT residences_user_id_fkey,
    ADD CONSTRAINT residences_user_id_fkey
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;
