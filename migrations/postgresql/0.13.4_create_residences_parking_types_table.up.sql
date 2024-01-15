CREATE TABLE IF NOT EXISTS residences_parking_types (
    residence_id INT REFERENCES residences (id),
    parking_type_id INT REFERENCES parking_types (id),
    PRIMARY KEY (residence_id, parking_type_id)
);