CREATE TABLE IF NOT EXISTS surat(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nomor_surat VARCHAR(50) NOT NULL,
    judul_surat VARCHAR(100) NOT NULL,
    tanggal_surat DATE,
    pengirim VARCHAR(100) NOT NULL,
    penerima VARCHAR(100) NOT NULL,
    perihal TEXT,
    user_id INT,
    jenis_surat ENUM('masuk', 'keluar', 'lainnya'),
    catatan TEXT,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);  

INSERT INTO surat VALUES(100, "291023/PM-128", "PENGIRIMAN MATERIAL", NULL, "PT. METALINDO", "PT. HIROTA TECH", "Pengiriman Material", NULL, "masuk", NULL, NULL);
INSERT INTO surat VALUES(101, "291023/PM-129", "PENGIRIMAN MATERIAL", NULL, "PT. METALINDO", "PT. HIROTA TECH", "Pengiriman Material", NULL, "masuk", NULL, NULL);
INSERT INTO surat VALUES(103, "291023/PM-130", "PENGIRIMAN MATERIAL", NULL, "PT. METALINDO", "PT. HIROTA TECH", "Pengiriman Material", NULL, "masuk", NULL, NULL);