CREATE TABLE IF NOT EXISTS product_variant_collections (
    id INT AUTO_INCREMENT PRIMARY KEY,
    variant_id INT NOT NULL,
    variant_key_id INT NOT NULL,
    variant_value_id INT NOT NULL,

    FOREIGN KEY (variant_id) REFERENCES product_variant(id) ON DELETE CASCADE,
    FOREIGN KEY (variant_key_id) REFERENCES product_variant_keys(id) ON DELETE CASCADE,
    FOREIGN KEY (variant_value_id) REFERENCES product_variant_values(id) ON DELETE CASCADE
);