CREATE TABLE IF NOT EXISTS product_variant (
    id INT AUTO_INCREMENT PRIMARY KEY,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);