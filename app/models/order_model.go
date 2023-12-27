package models

import (
	"context"
	"database/sql"
	"errors"
	"logistica/app/config"
	"logistica/app/utility"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	ID               sql.NullInt64  `gorm:"primaryKey" json:"id"`
	Buyer            sql.NullString `gorm:"column:buyer" json:"buyer"`
	NumberPhoneBuyer sql.NullString `gorm:"column:number_phone_buyer" json:"number_phone_buyer"`
	Receiver         sql.NullString `gorm:"column:receiver" json:"receiver"`
	ShippingAddress  sql.NullString `gorm:"column:shipping_address" json:"shipping_address"`
	Documentation    sql.NullByte   `gorm:"column:documentation" json:"documentation"`
	Pieces           sql.NullInt64  `gorm:"column:pieces" json:"pieces"`
	TotalPrice       sql.NullInt64  `gorm:"column:total_price" json:"total_price"`

	// Foreign Key
	Product   Product       `gorm:"foreignKey:ProductID" json:"product"`
	ProductID sql.NullInt64 `gorm:"column:product_id" json:"product_id"`
	Status    OrderStatus   `gorm:"foreignKey:StatusID" json:"order_status"`
	StatusID  sql.NullInt64 `gorm:"column:status_id" json:"status_id"`
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

func (o *Order) NewOrder(tx *sql.Tx, buyer, numberPhone, address string, pieces, totalPrice, productID, statusID int64) error {
	mutex.Lock()
	defer mutex.Unlock()

	var query string = "INSERT INTO orders(buyer, number_phone_buyer, shipping_address, pieces, total_price, product_id, status_id) VALUES(?, ?, ?, ?, ?, ?, ?);"
	_, err := tx.Exec(query, buyer, numberPhone, address, pieces, totalPrice, productID, statusID)
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
	// Kueri untuk melakukan JOIN TABLES
	var query string = `
	SELECT 
		o.id, 
		o.buyer, 
		o.number_phone_buyer, 
		o.receiver, 
		o.shipping_address, 
		o.documentation, 
		o.pieces, 
		o.total_price, 
		o.product_id, 
		o.status_id AS orders,
		p.name AS product_name,
		os.name AS status_name
	FROM 
		orders o
	JOIN 
		products p ON o.product_id = p.id
	JOIN
		order_statuses os ON o.status_id = os.id;
	`

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
			&order.Product.Name,
			&order.Status.Name,
		)

		if err != nil {
			return nil, err
		}

		orderMap := map[string]interface{}{
			"id":                 order.ID.Int64,
			"buyer":              order.Buyer.String,
			"number_phone_buyer": order.NumberPhoneBuyer.String,
			"receiver":           order.Receiver.String,
			"shipping_address":   order.ShippingAddress.String,
			"documentation":      order.Documentation,
			"pieces":             order.Pieces.Int64,
			"total_price":        utility.RupiahFormat(order.TotalPrice.Int64),
			"product_name":       utility.CapitalizeAll(order.Product.Name.String),
			"order_status":       order.Status.Name.String,
		}

		orders = append(orders, orderMap)
	}

	return orders, nil
}
