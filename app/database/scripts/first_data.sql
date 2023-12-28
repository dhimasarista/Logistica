-- Default Employee untuk default user
INSERT INTO employees VALUES(1, "administrator", "0", "0", NULL, 1, 1, NOW(), NOW(), NULL);
-- Default User as administrator
INSERT INTO users(id, username, password, employee_id) VALUES(1, "0x0002", "$2a$12$jb.qLEDHWmvFptryo8J/e.LnxhxNu9N5mmH.IEmHkjMvNYbb9f.iq", 1, NOW(), NOW(), NULL);
-- Positions
INSERT INTO positions VALUES(2222, "software engineer", NOW(), NOW(), NULL);
INSERT INTO positions VALUES(2223, "human resource", NOW(), NOW(), NULL);
INSERT INTO positions VALUES(2224, "admin staff", NOW(), NOW(), NULL);

INSERT INTO order_statuses VALUES(1, "on process", NOW(), NOW(), NULL);
INSERT INTO order_statuses VALUES(2, "on delivery", NOW(), NOW(), NULL);
INSERT INTO order_statuses VALUES(3, "received", NOW(), NOW(), NULL);
INSERT INTO order_statuses VALUES(4, "returned", NOW(), NOW(), NULL);
INSERT INTO order_statuses VALUES(5, "fail", NOW(), NOW(), NULL);
INSERT INTO order_statuses VALUES(6, "cancelled", NOW(), NOW(), NULL);