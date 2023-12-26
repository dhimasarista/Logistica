package models

import (
	"context"
	"database/sql"
	"errors"
	"logistica/app/config"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	ID         sql.NullInt64 `gorm:"primaryKey;column:id" json:"id"`
	Pieces     sql.NullInt64 `gorm:"column:pieces" json:"pieces"`
	TotalPrice sql.NullInt64 `gorm:"column:total_price" json:"total_price"`
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

func (o *Order) NewOrder(tx *sql.Tx, id, pieces, totalPrice, productID, statusID, detailID int64) error {
	mutex.Lock()
	defer mutex.Unlock()

	var query string = "INSERT INTO orders(id, pieces, total_price, product_id, status_id, detail_id) VALUES(?, ?, ?, ?, ?, ?);"
	_, err := tx.Exec(query, id, pieces, totalPrice, productID, statusID, detailID)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errors.New("race condition, id has been taken")
			}
		}
		return err
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
