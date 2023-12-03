package config

import (
	"database/sql"
	"log"
)

func ConnectDB() *sql.DB {
	// Membuat koneksi database dengan database pooling
	var dsn string = "root@tcp(localhost:3306)/logistica?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}

	// Set maximum open connections and maximum idle connections
	// db.SetMaxOpenConns(10) // Atur jumlah maksimum koneksi terbuka
	// db.SetMaxIdleConns(5)  // Atur jumlah maksimum koneksi yang tetap terbuka

	// Pastikan koneksi database ditutup saat selesai
	return db
}
