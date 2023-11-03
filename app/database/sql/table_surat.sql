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