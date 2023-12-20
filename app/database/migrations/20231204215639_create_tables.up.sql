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

CREATE TABLE earnings(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    amount_received BIGINT,
    product_name VARCHAR(255) NOT NULL,
    pieces INT,
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

