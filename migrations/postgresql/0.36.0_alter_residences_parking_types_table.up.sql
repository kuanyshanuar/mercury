ALTER TABLE residences_parking_types
    DROP CONSTRAINT residences_parking_types_parking_type_id_fkey,
    ADD CONSTRAINT residences_parking_types_parking_type_id_fkey
        FOREIGN KEY (parking_type_id)
            REFERENCES parking_types(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

ALTER TABLE residences_parking_types
    DROP CONSTRAINT residences_parking_types_residence_id_fkey,
    ADD CONSTRAINT residences_parking_types_residence_id_fkey
        FOREIGN KEY (residence_id)
            REFERENCES residences(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;
