CREATE TABLE IF NOT EXISTS cottages_wall_types(
    wall_type_id INT references wall_types(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (wall_type_id, cottage_id)
);

CREATE TABLE IF NOT EXISTS cottages_elevator_types(
    elevator_type_id INT references elevator_types(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (elevator_type_id, cottage_id)
);

CREATE TABLE IF NOT EXISTS cottages_heating_types(
    heating_type_id INT references heating_types(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (heating_type_id, cottage_id)
);

CREATE TABLE IF NOT EXISTS cottages_interior_decorations(
    interior_decoration_id INT references interior_decorations(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (interior_decoration_id, cottage_id)
);

CREATE TABLE IF NOT EXISTS cottages_parking_types(
    parking_type_id  INT references parking_types(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (parking_type_id, cottage_id)
);

CREATE TABLE IF NOT EXISTS cottages_warming_types(
    warming_type_id INT references warming_types(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (warming_type_id, cottage_id)
);

CREATE TABLE IF NOT EXISTS cottages_purchase_methods(
    purchase_method_id INT REFERENCES purchase_methods(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (purchase_method_id, cottage_id)
);

CREATE TABLE IF NOT EXISTS cottages_windows(
    window_type_id INT references window_types(id),
    cottage_id INT references cottages(id),
    PRIMARY KEY (window_type_id, cottage_id)
);