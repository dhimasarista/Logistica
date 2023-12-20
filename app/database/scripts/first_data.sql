-- Default Employee untuk default user
INSERT INTO employees VALUES(1, "administrator", "0", "0", NULL, 1, 1);
-- Default User as administrator
INSERT INTO users(id, username, password, employee_id) VALUES(1, "0x0002", "$2a$12$jb.qLEDHWmvFptryo8J/e.LnxhxNu9N5mmH.IEmHkjMvNYbb9f.iq", 1);
-- Positions
INSERT INTO positions VALUES(2222, "software engineer");
INSERT INTO positions VALUES(2223, "human resource");
INSERT INTO positions VALUES(2224, "admin staff");

INSERT INTO order_status VALUES(1, "on delivery")
INSERT INTO order_status VALUES(2, "received")
INSERT INTO order_status VALUES(3, "returned")
INSERT INTO order_status VALUES(4, "fail")