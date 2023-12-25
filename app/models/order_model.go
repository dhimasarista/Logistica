package models

import (
	"context"
	"database/sql"
	"logistica/app/config"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID     sql.NullInt64 `gorm:"primaryKey;column:id" json:"id"`
	Pieces sql.NullInt64 `gorm:"column:pieces" json:"pieces"`
	// Foreign Key
	ProductID sql.NullInt64 `gorm:"column:product_id" json:"product_id"`
	StatusID  sql.NullInt64 `gorm:"column:status_id" json:"status_id"`
	DetailID  sql.NullInt64 `gorm:"column:detail_id" json:"detail_id"`

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

func (o *Order) NewOrder(id, pieces, productID, statusID, detailID int64) error {
	var db = config.ConnectGormDB()

	results := db.Exec("INSERT INTO orders(id, pieces, product_id, status_id, detail_id) VALUES(?, ?, ?, ?, ?);", id, pieces, productID, statusID, detailID)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (o *Order) FindAll() ([]map[string]any, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var query string = "SELECT id, pieces, product_id, status_id, detail_id FROM orders"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var orders []map[string]any
	for rows.Next() {
		err := rows.Scan(
			&o.ID.Int64,
			&o.Pieces.Int64,
			&o.ProductID.Int64,
			&o.StatusID.Int64,
			&o.DetailID.Int64,
		)

		if err != nil {
			return nil, err
		}

		order := map[string]interface{}{
			"id":        o.ID.Int64,
			"pieces":    o.Pieces.Int64,
			"productId": o.ProductID.Int64,
			"statusId":  o.StatusID.Int64,
			"detailId":  o.ProductID.Int64,
		}

		orders = append(orders, order)
	}

	return orders, nil
}
