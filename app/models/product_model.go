package models

import (
	"context"
	"database/sql"
	"logistica/app/config"
	"time"
)

type Product struct {
	ID               sql.NullInt64  `json:"id"`
	Name             sql.NullString `json:"name"`
	SerialNumber     sql.NullString `json:"serial_number"`
	ManufacturerID   sql.NullInt64  `json:"manufacturer_id"`
	ManufacturerName sql.NullString `json:"manufacturer_name"`
	Stocks           sql.NullInt64  `json:"stocks"`
	Price            sql.NullInt64  `json:"price"`
	Weight           sql.NullInt64  `json:"weight"`
	CategoryID       sql.NullInt64  `json:"category_id"`
	CategoryName     sql.NullString `json:"category_name"`
}

func (p *Product) GetById(id int) error {
	var db = config.ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	query := `
	SELECT
		p.id, 
		p.name, 
		p.serial_number, 
		p.manufacturer_id, 
		p.stocks, 
		p.price, 
		p.weight, 
		p.category_id AS products, 
		m.name AS manufacturer_name, 
		c.name AS category_name 
	FROM 
		products p 
	JOIN 
		manufacturer m ON p.manufacturer_id = m.id 
	JOIN 
		product_category c ON p.category_id = c.id
	WHERE p.id = ?;
	`
	err := db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.SerialNumber,
		&p.ManufacturerID,
		&p.Stocks,
		&p.Price,
		&p.Weight,
		&p.CategoryID,
		&p.ManufacturerName,
		&p.CategoryName,
	)
	if err != nil {

		return err
	}

	return nil
}

func (p *Product) FindAll() error {
	var db = config.ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := "SELECT * FROM products"
	err := db.QueryRowContext(ctx, query)
	if err != nil {
		return err.Err()
	}

	return nil
}
