CREATE TABLE IF NOT EXISTS product_variant_warehouse (
    id SERIAL PRIMARY KEY,
    stock INT NOT NULL,
    warehouse_id INT NOT NULL,
    variant_id INT NOT NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE,
    FOREIGN KEY (variant_id) REFERENCES product_variants(id) ON DELETE CASCADE
);