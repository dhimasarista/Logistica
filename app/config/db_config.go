package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig menyimpan konfigurasi database
type DBConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	Name      string
	ParseTime bool
}

func generateDSN(config DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%t",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.ParseTime,
	)
}

// ReadDBConfig membaca konfigurasi database dari Viper
func readDBConfig() DBConfig {
	viper.SetConfigFile("../app/config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s", err)
	}

	var dbConfig DBConfig
	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
		fmt.Printf("Unable to decode database config: %s", err)
	}

	return dbConfig
}

var dbConfig DBConfig = readDBConfig()

// ConnectSQLDB membuat koneksi *sql.DB dan mengembalikannya
func ConnectSQLDB() *sql.DB {
	dsn := generateDSN(dbConfig)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

// ConnectGormDB membuat koneksi *gorm.DB dan mengembalikkannya
func ConnectGormDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(generateDSN(dbConfig)), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}
