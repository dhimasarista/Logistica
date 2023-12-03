CREATE TABLE products(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    serial_number VARCHAR(255),
    manufacturer_id int,
    FOREIGN KEY(manufacturer_id) REFERENCES manufacturer(id),
    stocks BIGINT,
    price INT,
    weight INT,
    category_id INT,
    FOREIGN KEY(category_id) REFERENCES product_category(id)
);