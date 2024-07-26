CREATE TABLE IF NOT EXISTS wishlist_variants (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    variant_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (variant_id) REFERENCES product_variants(id) ON DELETE CASCADE
);