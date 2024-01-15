DROP TABLE residences_interior_decorations_types;

CREATE TABLE residences_interior_decorations (
    residence_id             BIGINT REFERENCES residences (id),
    interior_decoration_id   BIGINT REFERENCES interior_decorations (id),
    PRIMARY KEY (residence_id, interior_decoration_id)
);