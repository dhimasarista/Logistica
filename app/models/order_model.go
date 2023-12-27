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
	ID               uint   `gorm:"primaryKey" json:"id"`
	Buyer            string `gorm:"column:buyer" json:"buyer"`
	NumberPhoneBuyer string `gorm:"column:number_phone_buyer" json:"number_phone_buyer"`
	Receiver         string `gorm:"column:receiver" json:"receiver"`
	ShippingAddress  string `gorm:"column:shipping_address" json:"shipping_address"`
	Documentation    []byte `gorm:"column:documentation" json:"documentation"`
	Pieces           int    `gorm:"column:pieces" json:"pieces"`
	TotalPrice       int    `gorm:"column:total_price" json:"total_price"`

	// Foreign Key
	ProductID uint `gorm:"column:product_id" json:"product_id"`
	StatusID  uint `gorm:"column:status_id" json:"status_id"`

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
func (o *Order) FindAll() ([]map[string]interface{}, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var query string = "SELECT id, buyer, number_phone_buyer, receiver, shipping_address, documentation, pieces, total_price, product_id, status_id FROM orders"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var orders []map[string]interface{}
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.Buyer,
			&order.NumberPhoneBuyer,
			&order.Receiver,
			&order.ShippingAddress,
			&order.Documentation,
			&order.Pieces,
			&order.TotalPrice,
			&order.ProductID,
			&order.StatusID,
		)

		if err != nil {
			return nil, err
		}

		orderMap := map[string]interface{}{
			"id":                 order.ID,
			"buyer":              order.Buyer,
			"number_phone_buyer": order.NumberPhoneBuyer,
			"receiver":           order.Receiver,
			"shipping_address":   order.ShippingAddress,
			"documentation":      order.Documentation,
			"pieces":             order.Pieces,
			"total_price":        order.TotalPrice,
			"product_id":         order.ProductID,
			"status_id":          order.StatusID,
		}

		orders = append(orders, orderMap)
	}

	return orders, nil
}
