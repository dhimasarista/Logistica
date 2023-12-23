package models

import (
	"database/sql"
	"logistica/app/config"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID     sql.NullInt64 `gorm:"primaryKey;column:id" json:"id"`
	Pieces sql.NullInt32 `gorm:"column:pieces" json:"pieces"`
	// Foreign Key
	ProductID sql.NullString `gorm:"column:product_id" json:"product_id"`
	StatusID  sql.NullString `gorm:"column:status_id" json:"status_id"`
	DetailID  sql.NullString `gorm:"column:detail_id" json:"detail_id"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (o *Order) TotalOrders() (int, error) {
	var db = config.ConnectGormDB()

	var total int
	results := db.Raw("SELECT COUNT(*) AS total FROM orders;").Scan(&total)
	if results.Error != nil {
		return -1, results.Error
	}

	return total, nil
}
