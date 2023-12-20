package config

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectSQLDB membuat koneksi *sql.DB dan mengembalikannya
func ConnectSQLDB() *sql.DB {
	dsn := "user_dev:vancouver@tcp(localhost:3306)/logistica?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

// ConnectGormDB membuat koneksi *gorm.DB dan mengembalikkannya
func ConnectGormDB() *gorm.DB {
	dsn := "user_dev:vancouver@tcp(localhost:3306)/logistica?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}
