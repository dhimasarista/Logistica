CREATE TABLE positions (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE employees(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    number_phone VARCHAR(32) NOT NULL,
    position_id INT,
    identity_card LONGBLOB NULL,
    is_user TINYINT(1) NOT NULL,
    is_superuser TINYINT(1) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (position_id) REFERENCES positions(id)
);

CREATE TABLE product_categories (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE manufacturers (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(255),
    password TEXT,
    employee_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);

CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    serial_number VARCHAR(255),
    manufacturer_id INT,
    stocks BIGINT,
    price INT,
    weight INT,
    category_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturers(id),
    FOREIGN KEY (category_id) REFERENCES product_categories(id)
);

CREATE TABLE stock_records (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    amount INT NOT NULL,
    before_record INT NULL,
    after_record INT NOT NULL,
    product_id INT,
    description TEXT,
    is_addition TINYINT(1) DEFAULT 1, -- 1 untuk penambahan, 0 untuk yang lain
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE order_statuses (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    buyer VARCHAR(255),
    number_phone_buyer VARCHAR(255),
    receiver VARCHAR(255) NULL,
    shipping_address VARCHAR(255),
    documentation LONGBLOB,
    pieces INT NOT NULL,
    total_price INT NOT NULL,
    product_id INT,
    status_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (status_id) REFERENCES order_statuses(id)
);

CREATE TABLE earnings (
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    amount_received BIGINT,
    product_name VARCHAR(255) NOT NULL,
    pieces INT,
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);