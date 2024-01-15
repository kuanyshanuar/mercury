CREATE TABLE IF NOT EXISTS residences_interior_decorations_types (
    residence_id INT REFERENCES residences (id),
    interior_decoration_type_id INT REFERENCES interior_decorations (id),
    PRIMARY KEY (residence_id, interior_decoration_type_id)
);