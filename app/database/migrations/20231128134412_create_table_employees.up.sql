CREATE TABLE employees(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(64) NOT NULL,
    address VARCHAR(255) NOT NULL,
    number_phone VARCHAR(32) NOT NULL,
    position VARCHAR(32) NOT NULL,
    is_user TINYINT NOT NULL,
    is_superuser TINYINT NOT NULL
) ENGINE = InnoDB;

-- INSERT INTO employees VALUES(1, "Administrator", "0", "0", "Administrator", 1, 1);