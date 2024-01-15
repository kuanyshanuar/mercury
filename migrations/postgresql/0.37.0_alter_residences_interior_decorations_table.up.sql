ALTER TABLE residences_interior_decorations
    DROP CONSTRAINT residences_interior_decorations_interior_decoration_id_fkey,
    ADD CONSTRAINT residences_interior_decorations_interior_decoration_id_fkey
        FOREIGN KEY (interior_decoration_id)
            REFERENCES interior_decorations(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

ALTER TABLE residences_interior_decorations
    DROP CONSTRAINT residences_interior_decorations_residence_id_fkey,
    ADD CONSTRAINT residences_interior_decorations_residence_id_fkey
        FOREIGN KEY (residence_id)
            REFERENCES residences(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;
