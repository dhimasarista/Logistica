package models

import (
	"database/sql"
	"errors"
	"logistica/app/config"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID            sql.NullInt64  `gorm:"primaryKey;column:id" json:"id"`
	Buyer         sql.NullString `gorm:"column:buyer" json:"buyer"`
	NumberPhone   sql.NullString `gorm:"column:number_phone_buyer" json:"number_phone_buyer"`
	Receiver      sql.NullString `gorm:"column:receiver" json:"receiver"`
	Address       sql.NullString `gorm:"column:shipping_address" json:"shipping_address"`
	Documentation sql.NullByte   `gorm:"column:documentation" json:"documentation"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (od *OrderDetail) GetByID(id int) error {
	var db = config.ConnectGormDB()

	query := "SELECT * FROM order_detail WHERE id = ?"
	result := db.Raw(query, id).Scan(&od)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (od *OrderDetail) NewOrder(buyer, numberPhone, address string) (sql.Result, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectSQLDB()
	defer db.Close()

	var query string = "INSERT INTO order_detail(buyer, number_phone_buyer, shipping_address) VALUES(?, ?, ?);"
	result, err := db.Exec(query, buyer, numberPhone, address)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil, errors.New("race condition, id has been taken")
			}
		}
		return result, err
	}

	return result, nil
}
