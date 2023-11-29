CREATE TABLE shipment(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    order_id INT,
    FOREIGN KEY(order_id) REFERENCES orders(id),
    delivery_number VARCHAR(255),
    shipping_address VARCHAR(255),
    coordinate VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);