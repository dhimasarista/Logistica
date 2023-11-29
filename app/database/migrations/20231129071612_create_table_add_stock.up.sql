CREATE TABLE add_stock(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    amount_added BIGINT NOT NULL,
    product_id INT,
    FOREIGN KEY (product_id) REFERENCES products(id)
);