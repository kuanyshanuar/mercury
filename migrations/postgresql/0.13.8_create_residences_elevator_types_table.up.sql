CREATE TABLE IF NOT EXISTS residences_elevator_types (
    residence_id INT REFERENCES residences (id),
    elevator_type_id INT REFERENCES elevator_types (id),
    PRIMARY KEY (residence_id, elevator_type_id)
);