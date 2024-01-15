ALTER TABLE residences_elevator_types
    DROP CONSTRAINT residences_elevator_types_elevator_type_id_fkey,
    ADD CONSTRAINT residences_elevator_types_elevator_type_id_fkey
        FOREIGN KEY (elevator_type_id)
            REFERENCES elevator_types(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

ALTER TABLE residences_elevator_types
    DROP CONSTRAINT residences_elevator_types_residence_id_fkey,
    ADD CONSTRAINT residences_elevator_types_residence_id_fkey
        FOREIGN KEY (residence_id)
            REFERENCES residences(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;
