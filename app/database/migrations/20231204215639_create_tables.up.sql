CREATE TABLE positions (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE employees(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    number_phone VARCHAR(32) NOT NULL,
    position_id INT,
    is_user TINYINT(1) NOT NULL,
    is_superuser TINYINT(1) NOT NULL,
    FOREIGN KEY (position_id) REFERENCES positions(id)
);

CREATE TABLE product_category(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255)
);

CREATE TABLE manufacturer(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255)
);

CREATE TABLE users(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(255),
    password TEXT, 
    employee_id INT,
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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

CREATE TABLE add_stock(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    amount_added BIGINT NOT NULL,
    product_id INT,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE order_status(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255)
);

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

CREATE TABLE order_detail(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    buyer VARCHAR(255),
    receiver VARCHAR(255),
    shipper VARCHAR(255),
    documentation LONGBLOB
);

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

CREATE TABLE shipment(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    order_id INT,
    FOREIGN KEY(order_id) REFERENCES orders(id),
    delivery_number VARCHAR(255),
    shipping_address VARCHAR(255),
    coordinate VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Input data ketika migrasi

-- Positions
INSERT INTO positions VALUES(2222, "software engineer");
INSERT INTO positions VALUES(2223, "human resource");
INSERT INTO positions VALUES(2224, "admin staff");
-- Default Employee untuk default user
INSERT INTO employees VALUES(1, "administrator", "0", "0", NULL, 1, 1);
-- Default User as administrator
INSERT INTO users(id, username, password, employee_id) VALUES(1, "0x0002", "$2a$12$jb.qLEDHWmvFptryo8J/e.LnxhxNu9N5mmH.IEmHkjMvNYbb9f.iq", 1);
-- Default Manufacturer
INSERT INTO manufacturer VALUES(9100, "amd");
INSERT INTO manufacturer VALUES(9110, "intel");
-- Default Product Category
INSERT INTO product_category VALUES (890, "processor");
INSERT INTO product_category VALUES (891, "ram");
INSERT INTO product_category VALUES (892, "motherboard");
-- Default Products
INSERT INTO products VALUES(1023, "ryzen 5500x", "AR47395349854365", 9100, 28, 1490000, 300, 890);
INSERT INTO products VALUES(1024, "i5 5400h", "IP834867242364832", 9100, 34, 1490000, 249, 890);