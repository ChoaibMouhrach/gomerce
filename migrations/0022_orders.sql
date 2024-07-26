CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    shipping_price FLOAT NOT NULL,
    user_id INT NOT NULL,
    status_id INT NOT NULL,
    payment_option_id INT NOT NULL,
    shipping_option_id INT NOT NULL,
    shipping_address_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (status_id) REFERENCES order_statuses(id) ON DELETE CASCADE,
    FOREIGN KEY (shipping_option_id) REFERENCES shipping_options(id) ON DELETE CASCADE,
    FOREIGN KEY (payment_option_id) REFERENCES payment_options(id) ON DELETE CASCADE,
    FOREIGN KEY (shipping_address_id) REFERENCES shipping_addresses(id) ON DELETE CASCADE
);