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