CREATE TABLE IF NOT EXISTS residences_heating_types (
    residence_id INT REFERENCES residences (id),
    heating_type_id INT REFERENCES heating_types (id),
    PRIMARY KEY (residence_id, heating_type_id)
);