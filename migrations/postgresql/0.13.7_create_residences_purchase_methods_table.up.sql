CREATE TABLE IF NOT EXISTS residences_purchase_methods (
    residence_id INT REFERENCES residences (id),
    purchase_method_id INT REFERENCES purchase_methods (id),
    PRIMARY KEY (residence_id, purchase_method_id)
);