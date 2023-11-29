CREATE TABLE order_detail(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    buyer VARCHAR(255),
    receiver VARCHAR(255),
    shipper VARCHAR(255),
    documentation LONGBLOB
);