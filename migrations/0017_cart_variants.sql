CREATE TABLE IF NOT EXISTS cart_variants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    quantity INT NOT NULL,
    cart_id INT NOT NULL,
    variant_id INT NOT NULL,
    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    FOREIGN KEY (variant_id) REFERENCES product_variants(id) ON DELETE CASCADE
);