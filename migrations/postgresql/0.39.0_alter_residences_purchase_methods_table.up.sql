ALTER TABLE residences_purchase_methods
    DROP CONSTRAINT residences_purchase_methods_purchase_method_id_fkey,
    ADD CONSTRAINT residences_purchase_methods_purchase_method_id_fkey
        FOREIGN KEY (purchase_method_id)
            REFERENCES purchase_methods(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

ALTER TABLE residences_purchase_methods
    DROP CONSTRAINT residences_purchase_methods_residence_id_fkey,
    ADD CONSTRAINT residences_purchase_methods_residence_id_fkey
        FOREIGN KEY (residence_id)
            REFERENCES residences(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;
