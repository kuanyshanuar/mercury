-- Cottage wall types constraint
ALTER TABLE cottages_wall_types
    DROP CONSTRAINT cottages_wall_types_wall_type_id_fkey,
    ADD CONSTRAINT cottages_wall_types_wall_type_id_fkey
        FOREIGN KEY (wall_type_id)
            REFERENCES wall_types(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_wall_types
    DROP CONSTRAINT cottages_wall_types_cottage_id_fkey,
    ADD CONSTRAINT cottages_wall_types_cottage_id_fkey
        FOREIGN KEY (cottage_id)
            REFERENCES cottages(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

-- Cottage heating types constraint
ALTER TABLE cottages_heating_types
    DROP CONSTRAINT cottages_heating_types_heating_type_id_fkey,
    ADD CONSTRAINT cottages_heating_types_heating_type_id_fkey
        FOREIGN KEY (heating_type_id)
            REFERENCES heating_types(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_heating_types
    DROP CONSTRAINT cottages_heating_types_cottage_id_fkey,
    ADD CONSTRAINT cottages_heating_types_cottage_id_fkey
        FOREIGN KEY (cottage_id)
            REFERENCES cottages(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_interior_decorations
    DROP CONSTRAINT cottages_interior_decorations_interior_decoration_id_fkey,
    ADD CONSTRAINT cottages_interior_decorations_interior_decoration_id_fkey
        FOREIGN KEY (interior_decoration_id)
            REFERENCES interior_decorations(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_interior_decorations
    DROP CONSTRAINT cottages_interior_decorations_cottage_id_fkey,
    ADD CONSTRAINT cottages_interior_decorations_cottage_id_fkey
        FOREIGN KEY (cottage_id)
            REFERENCES cottages(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

-- Cottage cottages_parking_types constraint
ALTER TABLE cottages_parking_types
    DROP CONSTRAINT cottages_parking_types_parking_type_id_fkey,
    ADD CONSTRAINT cottages_parking_types_parking_type_id_fkey
        FOREIGN KEY (parking_type_id)
            REFERENCES parking_types(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_parking_types
    DROP CONSTRAINT cottages_parking_types_cottage_id_fkey,
    ADD CONSTRAINT cottages_parking_types_cottage_id_fkey
        FOREIGN KEY (cottage_id)
            REFERENCES cottages(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

-- Cottage cottages_warming_types constraint
ALTER TABLE cottages_warming_types
    DROP CONSTRAINT cottages_warming_types_warming_type_id_fkey,
    ADD CONSTRAINT cottages_warming_types_warming_type_id_fkey
        FOREIGN KEY (warming_type_id)
            REFERENCES warming_types(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_warming_types
    DROP CONSTRAINT cottages_warming_types_cottage_id_fkey,
    ADD CONSTRAINT cottages_warming_types_cottage_id_fkey
        FOREIGN KEY (cottage_id)
            REFERENCES cottages(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

-- Cottage cottages_purchase_methods constraint
ALTER TABLE cottages_purchase_methods
    DROP CONSTRAINT cottages_purchase_methods_purchase_method_id_fkey,
    ADD CONSTRAINT cottages_purchase_methods_purchase_method_id_fkey
        FOREIGN KEY (purchase_method_id)
            REFERENCES purchase_methods(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_purchase_methods
    DROP CONSTRAINT cottages_purchase_methods_cottage_id_fkey,
    ADD CONSTRAINT cottages_purchase_methods_cottage_id_fkey
        FOREIGN KEY (cottage_id)
            REFERENCES cottages(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

-- Cottage cottages_windows constraint
ALTER TABLE cottages_windows
    DROP CONSTRAINT cottages_windows_window_type_id_fkey,
    ADD CONSTRAINT cottages_windows_window_type_id_fkey
        FOREIGN KEY (window_type_id)
            REFERENCES window_types(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

ALTER TABLE cottages_windows
    DROP CONSTRAINT cottages_windows_cottage_id_fkey,
    ADD CONSTRAINT cottages_windows_cottage_id_fkey
        FOREIGN KEY (cottage_id)
            REFERENCES cottages(id)
            ON UPDATE CASCADE
            ON DELETE SET NULL;

