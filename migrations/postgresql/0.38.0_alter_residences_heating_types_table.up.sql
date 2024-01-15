ALTER TABLE residences_heating_types
    DROP CONSTRAINT residences_heating_types_heating_type_id_fkey,
    ADD CONSTRAINT residences_heating_types_heating_type_id_fkey
        FOREIGN KEY (heating_type_id)
            REFERENCES heating_types(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

ALTER TABLE residences_heating_types
    DROP CONSTRAINT residences_heating_types_residence_id_fkey,
    ADD CONSTRAINT residences_heating_types_residence_id_fkey
        FOREIGN KEY (residence_id)
            REFERENCES residences(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;
