CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    read BOOLEAN NOT NULL DEFAULT FALSE,
    user_id INT,
    type_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (type_id) REFERENCES notification_types(id) ON DELETE CASCADE
);