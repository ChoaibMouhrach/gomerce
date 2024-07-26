CREATE TABLE IF NOT EXISTS product_variant_values (
    id SERIAL PRIMARY KEY,
    value TEXT NOT NULL,
    variant_key_id INT NOT NULL,
    FOREIGN KEY (variant_key_id) REFERENCES product_variant_keys(id) ON DELETE CASCADE
);