CREATE TABLE orders(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    pieces INT NOT NULL,
    product_id INT,
    status_id INT,
    detail_id INT,
    FOREIGN KEY(product_id) REFERENCES products(id),
    FOREIGN KEY(status_id) REFERENCES order_status(id),
    FOREIGN KEY(detail_id) REFERENCES order_detail(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);