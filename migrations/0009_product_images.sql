CREATE TABLE IF NOT EXISTS product_images (
    id SERIAL PRIMARY KEY,
    key TEXT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);