ALTER TABLE cottages
    DROP CONSTRAINT cottages_city_id_fkey,
    ADD CONSTRAINT cottages_city_id_fkey
        FOREIGN KEY (city_id)
            REFERENCES cities(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages
    DROP CONSTRAINT cottages_district_id_fkey,
    ADD CONSTRAINT cottages_district_id_fkey
        FOREIGN KEY (district_id)
            REFERENCES districts(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages
    DROP CONSTRAINT cottages_housing_class_id_fkey,
    ADD CONSTRAINT cottages_housing_class_id_fkey
        FOREIGN KEY (housing_class_id)
            REFERENCES housing_classes(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages
    DROP CONSTRAINT cottages_status_id_fkey,
    ADD CONSTRAINT cottages_status_id_fkey
        FOREIGN KEY (status_id)
            REFERENCES statuses(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages
    DROP CONSTRAINT cottages_user_id_fkey,
    ADD CONSTRAINT cottages_user_id_fkey
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;