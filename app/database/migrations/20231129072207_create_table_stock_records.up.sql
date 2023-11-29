CREATE TABLE stock_records(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    amount INT NOT NULL,
    before_record INT NULL,
    after_record INT NOT NULL,
    product_id INT,
    FOREIGN KEY(product_id) REFERENCES products(id),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);