package models

import (
	"database/sql"
	"logistica/app/config"
	"time"

	"gorm.io/gorm"
)

type Earning struct {
	ID   sql.NullInt64  `gorm:"primaryKey;column:id" json:"id"`
	Name sql.NullString `gorm:"column:name" json:"name"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (e *Earning) TotalEarnings() (int, error) {
	var db = config.ConnectGormDB()

	var query = "SELECT COALESCE(SUM(amount_received), 0) FROM earnings;"
	var total int
	result := db.Raw(query).Scan(&total)
	if result.Error != nil {
		return -1, result.Error
	}

	return total, nil
}

func (e *Earning) NewOrder(tx *sql.Tx, amountReceived int, productName string, pieces, price int) error {
	var query = "INSERT INTO earnings(amount_received, product_name, pieces, price) VALUES(?,?,?,?);"
	_, err := tx.Exec(query, amountReceived, productName, pieces, price)
	if err != nil {
		return err
	}

	return nil
}
